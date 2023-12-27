# Live Video Streaming Service 기술 스택 정리
- `실시간 인터넷 방송`에 사용될만한 기술 스택 내용 정리

## 연관 개념 정리
### 1. Codec  
- 데이터 스트림을 En/Decoding 하는 기술, 소프트웨어  
- 효율적인 데이터 전송 및 보관을 위해 데이터 열화를 감안한 손실 방식(MP3, JPEG, H.264 등)과 원본의 품질을 해치지 않기 위한 비손실 방식(FLAC, PNG, HuffYUV 등)이 있음  

### 2. MPEG(Moving Picture Experts Group)  
- ISO, IEC 산하의 멀티미디어 표준 개발 담당 그룹  
- 비디오 데이터에서 많이 사용되는 H.264(MPEG-4 AVC)와 같은 표준을 정의함  

### 3. FFMPEG  
- 비디오, 오디오의 En/Decoding, Transcoding 등 멀티미디어 파일 처리를 위한 오픈 소스 프로젝트이자 툴  
- 무료로 제공되며 멀티 플랫폼 환경에서 이용 가능

### 4. RTMP (Real Time Message Protocol)  
- 2009년 발표된 Adobe Flash에 이용하기 위한 오디오, 비디오 데이터 실시간 통신 프로토콜  
- TCP/IP 기반이며 많은 스트리밍 서비스에서 라이브 서비스를 위해 이용되고 있음  
- 2020년 Adobe Flash가 공식 지원 종료되며 직접 이용자에게 데이터를 전달하는 용도로는 사용되지 않는 추세  

### 5. HLS (HTTP Live Streaming)  
![HLS](https://docs-assets.developer.apple.com/published/f089b49e80af12371bab35ee7275c735/http-live-streaming-1~dark@2x.png)  
- 2009년 Apple에서 발표한 HTTP 기반 적응 비트레이트 스트리밍 통신 프로토콜  
- 비디오 데이터를 `.m3u8` 확장자로 된 재생목록을 기록한 파일과 `.ts` 확장자의 여러 개의 멀티미디어 파일로 분할하여 스트리밍에 이용  
- Apple 제품에서 호환되는 유일한 스트리밍 포로토콜이며 기타 환경에서는 HLS Player를 통해 재생 가능  
- 전송에 이용할 비디오 데이터의 경우 H.264, H.265 Encoding 이용 필요 ([참조](https://www.cloudflare.com/ko-kr/learning/video/what-is-mpeg-dash/))

### 6. DASH (Dynamic Adaptive Streaming over HTTP)  
- 2014년 MPEG에서 정의한 국제 표준 HTTP 기반 적응 비트레이트 스트리밍 통신 프로토콜  
- 비디오 데이터를 `.mpd` 확장자의 스트리밍 메타데이터 파일과 여러 개의 `.m4a`, `m4v` 파일로 분할하여 스트리밍에 이용  
  (`.mpd`파일이 HLS의 `.m3u8`의 역할 / `.m4a` 파일은 AAC Codec의 음성 파일 / `m4v` 파일은 H.264 Codec의 비디오 파일)
- 국제 표준인만큼 HLS와 달리 다양한 인코딩에 대응할 수 있음  

### 7. CDN (Content Delivery Network)  
- 효율적인 컨텐츠(데이터) 전송을 위해 캐시 서버를 이용하는 기술  
- 주로 물리적으로 먼 거리의 사용자에게 서비스를 제공하기 위해 사용

## 구현해본 부분
1. HLS Player : [hls.html](https://github.com/HashCitrine/testLiveVideoStreamingService/blob/master/page/hls.html)에 `hls.js` 예제를 이용하여 구성
2. Encoder : `FFMPEG`를 이용해 RTMP를 통해 전달 받은 데이터를 Encoding하는 서비스 생성  
   (HLS : [go](https://github.com/HashCitrine/testLiveVideoStreamingService/blob/master/go/service/hls.go#L31), [spring](https://github.com/HashCitrine/testLiveVideoStreamingService/blob/master/spring/src/main/java/com/example/hls/service/HlsService.java#L30), [nginx](https://github.com/HashCitrine/testLiveVideoStreamingService/blob/master/nginx/nginx.conf#L47) / DASH : [nginx](https://github.com/HashCitrine/testLiveVideoStreamingService/blob/master/nginx/nginx.conf#L57))
3. HLS : Encoder를 통해 생성된 파일을 HLS Player로 전달하여 재생 가능 여부 테스트([go](https://github.com/HashCitrine/testLiveVideoStreamingService/blob/master/go/handle/handle.go#L11), [spring](https://github.com/HashCitrine/testLiveVideoStreamingService/blob/master/spring/src/main/java/com/example/hls/service/HlsService.java#L17))
4. RTMP Server : `Nginx RTMP` Plugin을 이용하여 RTMP 서버([참조](https://github.com/HashCitrine/testLiveVideoStreamingService/blob/master/nginx/nginx.conf#L5))를 띄우고 `OBS`와 연결 가능 여부 테스트
5. Dash Player : [dash.html](https://github.com/HashCitrine/testLiveVideoStreamingService/blob/master/page/dash.html)에 `dash.js` 예제를 이용하여 구성

## 예상 서비스 논리 구성
![img](https://github.com/HashCitrine/testLiveVideoStreamingService/assets/38382859/f0ff3a77-fe00-4f83-b688-7e787f603ee7)  
실시간 스트리밍 서비스에 필요할 것으로 예상되는 개념을 정리하여 예상해본 논리적인 서비스 구성 예상도  

실제 서비스는 Web 및 Mobile UI, UI API 등도 포함이 되어 있을 것으로 생각되며  
운영상의 이유 혹은 분산 처리의 목적으로 각 서비스들이 더욱 세부적인 단위로 분리되어 있을 수 있으나  

`실시간 인터넷 방송`에 초점을 맞춰 간략하게 구성함

1. `Web Service` : 실시간 영상 수신 및 송출에 필요한 웹 서비스 (트래픽 부하 분산을 위한 캐시 서버 구성이 필요할 수 있음)
   - Streamer로부터 RTMP 요청을 통해 영상 수신
   - 영상을 Viewer에게 HLS로 송출
   - 수신 및 송출을 담당하는 서비스는 분리하여 각자 구성할 수 있음
2. `Middleware` : 시스템 전반을 관리하는 서비스
   - 각 서비스간 처리 연계 Interface(Message Queue나 In-Memory DB 등을 이용할 수 있음)
   - Logging, Monitoring 서비스와 연계하여 시스템 현황 기록 및 관리(ELK 등을 이용할 수 있음)
3. `Encoder` : 영상 데이터 Encoding 서비스
   - RTMP로 수신 받은 영상 데이터를 HLS 및 DASH 응답에 이용할 수 있는 형태로 Encoding (HLS : `.3m8u + .ts` / DASH : `.m4a + .m4v`)
   - 영상 데이터를 VOD 서비스용 Codec으로 Encoding
   - Encoding 처리된 영상은 NFS나 Object Storage 등에 저장하여 데이터를 이용하는 서비스에서 접근하는 방식으로 처리할 수 있음  
4. `Broadcast Service` : 방송 진행과 시청에 필요한 요청 처리
   - Streamer가 방송 정보 변경 시 Viewer의 화면에도 반영
   - 채팅 서비스
   - Viewer가 구독, 후원 시 해당 정보를 Streamer의 방송 시스템에 전달


## P2P (Peer-to-Peer)
- 서비스 공급자(중앙)로부터 데이터를 제공 받는 것이 아닌, 서비스 이용자(Peer)끼리 서로 통신하여 필요한 데이터를 주고 받는 형태
- 국내에서 스트리밍 서비스 시 네트워크 이용을 분산하기 위해 `P2P` 시스템을 대부분 도입하고 있음
- 주의사항 
  1) Peer 간의 데이터 공유를 위해 서로의 위치(IP)가 식별되어야 함 (보안 문제)
  2) Peer의 수가 적은 경우, 효용성이 떨어짐

### BitTorrent
![BitTorrent](https://camo.githubusercontent.com/d579cc5713418331221acb55180fb7068f9e39a762518ba63a9eb0378525440c/687474703a2f2f63646e2e6f7265696c6c792e636f6d2f65786365727074732f393738303539363531343433332f773264705f303330372e706e67)
- 가장 대표적인 P2P 프로토콜이자 동명의 서비스  
- 공유할 데이터를 조각(Piece)으로 나누어 Peer에게 공유
- Seeder : 데이터의 모든 조각을 가진 Peer
- Leecher : 데이터의 일부 조각을 가진 Peer
- Tracker : 파일 공유를 일으키는 중앙 서버로, 공유 대상인 Peer 목록과 공유 파일 고유 식별자인 Hash 값으로 구성된 정보(`Swarm`)를 관리함

### IPFS (InterPlanetary File System)
![IPFS Stack](https://camo.githubusercontent.com/0c6475ffd1e72afe459e43b220ebc242eceb715d1ead148f93d32825a87ab6a5/68747470733a2f2f696d6167652e736c696465736861726563646e2e636f6d2f756e7469746c65642d3136303331343132343630322f39352f646174612d737472756374757265732d696e2d616e642d6f6e2d697066732d32382d3633382e6a70673f63623d31343537393539363638)
- 분산형 파일 시스템에 데이터를 저장하고 네트워크를 통해 공유하기 위해 고안된 P2P용 프로토콜
- 기반 기술
  1. `DHT` (Distributed Hash Tables) : 네트워크에 참여한 노드들이 `Hash Table`을 각자 관리하여, 중앙 서버 없이 P2P 네트워크 형성
  2. `BitTorrent` : P2P 파일 공유 프로토콜
  3. `Merkle DAG` (Merkle Directed Acyclic Graphs) : IPFS에서 이용되는 데이터에 적용되는 자료 구조
     ![Merkle DAG](https://camo.githubusercontent.com/875eaff5b9a107a82f4ffe94cd591b06d36ccd7de8a43bc37c413813ef1270ff/687474703a2f2f77686174646f65737468657175616e747361792e636f6d2f6173736574732f696d616765732f697066735f6f626a656374735f6469726563746f72795f7374727563747572652e706e67)
     - Merkle Tree에서 DAG로 확장된 자료 구조로, Leaf(말단) 노드만 실제 데이터를 가지고 있던 Tree 구조에서 모든 노드들이 데이터를 가지는 형태가 되었음
     - 부모 노드가 자식 노드의 해시값으로 구성되는 점으로 인해, 아래의 특징을 가짐  
       1) 위변조가 불가(위변조 시 Root 노드의 해시값이 변경됨)
       2) 모든 노드가 사실상 연결된 구조(해시값)
       3) 자체적으로 무결성 확인 가능(Multi hash Checksum)
       4) 같은 데이터는 같은 해시값을 가지므로 데이터 중복 불가
  4. `SFS` (Self-certified FileSystems) : IPFS의 name system인 `IPNS`의 기반 기술
     ![SFS](https://camo.githubusercontent.com/90b06021ed07661d3d02777712ec5c2df2dd11fdd72d76c1d295a746847a36d7/68747470733a2f2f696d6167652e736c696465736861726563646e2e636f6d2f756e7469746c65642d3136303331343132343630322f39352f646174612d737472756374757265732d696e2d616e642d6f6e2d697066732d35392d3633382e6a70673f63623d31343537393539363638)
        - IPFS를 이용하기 위한 주소 자체에 `공개키로 해시화된 값`을 사용하여, 자체적으로 공유 위치(서버) 식별이 가능

## 참조
- (Apple) [HTTP 라이브 스트리밍](https://developer.apple.com/documentation/http-live-streaming#Encode-and-deliver-streaming-media)
- (Youtube) [YouTube Live Streaming API](https://developers.google.com/youtube/v3/live/life-of-a-broadcast?hl=ko)
- (아프리카TV) [Afreeca TV Open API](https://developers.afreecatv.com/?szWork=openapi)
- (Twitch) [Twitch API](https://dev.twitch.tv/docs/api)
- [Introduction to HLS](https://medium.com/@hongseongho/introduction-to-hls-e7186f411a02)
- [RTMP Streaming: 실시간 메시징 프로토콜에 대하여](https://growthvalue.tistory.com/178)
- [웹 소켓으로 실시간 데이터 전송하기](https://velog.io/@skh9797/%EC%9B%B9-%EC%86%8C%EC%BC%93%EC%9C%BC%EB%A1%9C-%EC%8B%A4%EC%8B%9C%EA%B0%84-%EB%8D%B0%EC%9D%B4%ED%84%B0-%EC%A0%84%EC%86%A1%ED%95%98%EA%B8%B0)
- [[방장기강] 방송장비 기술강좌 - 비디오 프로토콜](https://youtu.be/sUtIxxTkpOA?si=YjPP8R-ICrJ1hQvi)
- [nginx rtmp를 이용해서 실시간 스트리밍 구현 예제](https://qteveryday.tistory.com/372)
- [MPEG-DASH란 무엇입니까? | HLS와 DASH의 비교](https://www.cloudflare.com/ko-kr/learning/video/what-is-mpeg-dash/)
- [P2P의 개념](https://ddongwon.tistory.com/75)
- [BitTorrent (토렌트) 의 원리와 구조](https://blog.naver.com/manhdh/220038243469)
- [IPFS(InterPlanetary File System)이해하기](https://medium.com/@kblockresearch/8-ipfs-interplanetary-file-system-%EC%9D%B4%ED%95%B4%ED%95%98%EA%B8%B0-1%EB%B6%80-http-web%EC%9D%84-%EB%84%98%EC%96%B4%EC%84%9C-ipfs-web%EC%9C%BC%EB%A1%9C-46382a2a6539)