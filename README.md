# Live Video Streaming Service 구성해보기
- `실시간 인터넷 방송` 서비스 구성해보기

## 연관 개념 정리
1. Codec  
데이터 스트림을 En/Decoding 하는 기술, 소프트웨어  
효율적인 데이터 전송 및 보관을 위해 데이터 열화를 감안한 손실 방식(MP3, JPEG, H.264 등)과 원본의 품질을 해치지 않기 위한 비손실 방식(FLAC, PNG, HuffYUV 등)이 있음  

2. MPEG(Moving Picture Experts Group)  
ISO, IEC 산하의 멀티미디어 표준 개발 담당 그룹  
비디오 데이터에서 많이 사용되는 H.264(MPEG-4 AVC)와 같은 표준을 정의함  

3. FFMPEG  
비디오, 오디오의 En/Decoding, Transcoding 등 멀티미디어 파일 처리를 위한 오픈 소스 프로젝트이자 툴  
무료로 제공되며 멀티 플랫폼 환경에서 이용 가능

4. RTMP (Real Time Message Protocol)  
2009년 발표된 Adobe Flash에 이용하기 위한 오디오, 비디오 데이터 실시간 통신 프로토콜  
TCP/IP 기반이며 많은 스트리밍 서비스에서 라이브 서비스를 위해 이용되고 있음  
2020년 Adobe Flash가 공식 지원 종료되며 직접 이용자에게 데이터를 전달하는 용도로는 사용되지 않는 추세  

5. HLS (HTTP Live Streaming)  
![HLS](https://docs-assets.developer.apple.com/published/f089b49e80af12371bab35ee7275c735/http-live-streaming-1~dark@2x.png)  
2009년 Apple에서 발표한 HTTP 기반 적응 비트레이트 스트리밍 통신 프로토콜  
비디오 데이터를 `.m3u8` 확장자로 된 재생목록을 기록한 파일과 `.ts` 확장자의 여러 개의 멀티미디어 파일로 분할하여 스트리밍에 이용  
Apple 제품에서 호환되는 유일한 스트리밍 포로토콜이며 기타 환경에서는 HLS Player를 통해 재생 가능  
전송에 이용할 비디오 데이터의 경우 H.264, H.265 Encoding 이용 필요 ([참조](https://www.cloudflare.com/ko-kr/learning/video/what-is-mpeg-dash/))

6. DASH (Dynamic Adaptive Streaming over HTTP)  
2014년 MPEG에서 정의한 국제 표준 HTTP 기반 적응 비트레이트 스트리밍 통신 프로토콜  
국제 표준인만큼 HLS와 달리 다양한 인코딩에 대응할 수 있음  

7. CDN (Content Delivery Network)  
효율적인 컨텐츠(데이터) 전송을 위해 캐시 서버를 이용하는 기술  
주로 물리적으로 먼 거리의 사용자에게 서비스를 제공하기 위해 사용

## 구현해본 부분
1. HLS Player : [index.html](https://github.com/HashCitrine/testLiveVideoStreamingService/blob/master/page/index.html)에 `hls.js` 예제를 이용하여 구성
2. Encoder : `FFMPEG`를 이용해  `.3m8u` 및 `.ts` 파일로 Encoding하는 서비스 생성([go](https://github.com/HashCitrine/testLiveVideoStreamingService/blob/master/go/service/hls.go#L31), [spring](https://github.com/HashCitrine/testLiveVideoStreamingService/blob/master/spring/src/main/java/com/example/hls/service/HlsService.java#L30), [nginx](https://github.com/HashCitrine/testLiveVideoStreamingService/blob/master/nginx/nginx.conf#L47))
3. HLS : Encoder를 통해 생성된 파일을 HLS Player로 전달하여 재생 가능 여부 테스트([go](https://github.com/HashCitrine/testLiveVideoStreamingService/blob/master/go/handle/handle.go#L11), [spring](https://github.com/HashCitrine/testLiveVideoStreamingService/blob/master/spring/src/main/java/com/example/hls/service/HlsService.java#L17))
4. RTMP Server : `Nginx RTMP` Plugin을 이용하여 RTMP 서버([참조](https://github.com/HashCitrine/testLiveVideoStreamingService/blob/master/nginx/nginx.conf#L5))를 띄우고 `OBS`와 연결 가능 여부 테스트

## 예상 서비스 논리 구성
![img](https://github.com/HashCitrine/testLiveVideoStreamingService/assets/38382859/f0ff3a77-fe00-4f83-b688-7e787f603ee7)  
실시간 스트리밍 서비스에 필요할 것으로 예상되는 개념을 정리하여 예상해본 논리적인 서비스 구성 예상도  

실제 서비스는 Web 및 Mobile UI, UI API 등도 포함이 되어 있을 것으로 생각되며  
운영상의 이유 혹은 분산 처리의 목적으로 각 서비스들이 더욱 세부적인 단위로 분리되어 있을 수 있으나  

`실시간 인터넷 방송`에 초점을 맞춰 간략하게 구성함

1. `Web Service` : 실시간 영상 수신 및 송출에 필요한 웹 서비스 (CDN에 부합하는 캐시 구성이 필요할 수 있음)
   - Streamer로부터 RTMP 요청을 통해 영상 수신
   - 영상을 Viewer에게 HLS로 송출
   - 수신 및 송출을 담당하는 서비스는 분리하여 각자 구성할 수 있음
2. `Middleware` : 시스템 전반을 관리하는 서비스
   - 각 서비스간 처리 연계 Interface(Message Queue나 In-Memory DB 등을 이용할 수 있음)
   - Logging, Monitoring 서비스와 연계하여 시스템 현황 기록 및 관리(ELK 등을 이용할 수 있음)
3. `Encoder` : 영상 데이터 Encoding 서비스
   - RTMP로 수신 받은 영상 데이터를 HLS 응답에 이용할 `.3m8u` 및 `.ts` 파일로 Encoding
   - 영상 데이터를 VOD 서비스용 Codec으로 Encoding
   - Encoding 처리된 영상은 NFS나 Object Storage 등에 저장하여 데이터를 이용하는 서비스에서 접근하는 방식으로 처리할 수 있음  
4. `Broadcast Service` : 방송 진행과 시청에 필요한 요청 처리
   - Streamer가 방송 정보 변경 시 Viewer의 화면에도 반영
   - 채팅 서비스
   - Viewer가 구독, 후원 시 해당 정보를 Streamer의 방송 시스템에 전달

## 궁금점
1. 스트리밍 서비스 시 많은 시청자가 방송을 시청하는 경우 발생하는 물리적인 부하는 어떻게 분산할까? (CDN? Object Storage?)
2. 다시보기 영상은 언제 VOD용 인코딩 처리가 이루어질까? (방송 진행과 동시에? 방송 종료 이후에?)
3. 시청자들간 스트리밍 데이터를 공유하는 시스템이 데이터를 공유할 사용자들과 공유 받을 사용자들을 어떻게 선택하고 식별할까?  
   그리고 각 사용자들 환경에서 어떻게 통신할 서로 통신할까? (RPC?)

## 참조
- (Apple) [HTTP 라이브 스트리밍](https://developer.apple.com/documentation/http-live-streaming#Encode-and-deliver-streaming-media)
- (Youtube) [YouTube Live Streaming API](https://developers.google.com/youtube/v3/live/life-of-a-broadcast?hl=ko)
- (아프리카TV) [Afreeca TV Open API](https://developers.afreecatv.com/?szWork=openapi)
- (Twitch) [Twitch API](https://dev.twitch.tv/docs/api)
- [Introduction to HLS](https://medium.com/@hongseongho/introduction-to-hls-e7186f411a02)
- [RTMP Streaming: 실시간 메시징 프로토콜에 대하여](https://growthvalue.tistory.com/178)
- [단계별: SRT(Secure Reliable Transport) 라이브 이벤트 스트리밍](https://ko.studio.support.brightcove.com/live/get-started/step-step-live-srt.html)
- [웹 소켓으로 실시간 데이터 전송하기](https://velog.io/@skh9797/%EC%9B%B9-%EC%86%8C%EC%BC%93%EC%9C%BC%EB%A1%9C-%EC%8B%A4%EC%8B%9C%EA%B0%84-%EB%8D%B0%EC%9D%B4%ED%84%B0-%EC%A0%84%EC%86%A1%ED%95%98%EA%B8%B0)
- [[방장기강] 방송장비 기술강좌 - 비디오 프로토콜](https://youtu.be/sUtIxxTkpOA?si=YjPP8R-ICrJ1hQvi)
- [nginx rtmp를 이용해서 실시간 스트리밍 구현 예제](https://qteveryday.tistory.com/372)
- [MPEG-DASH란 무엇입니까? | HLS와 DASH의 비교](https://www.cloudflare.com/ko-kr/learning/video/what-is-mpeg-dash/)
- [P2P의 개념](https://ddongwon.tistory.com/75)
