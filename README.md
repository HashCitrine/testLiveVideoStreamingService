# Live Video Streaming Service 구성해보기
- `RTMP` & `HLS` Protocol Server를 구성하여 `실시간 인터넷 방송` 서비스 구성해보기

## 개념 정리
1. RTMP (Real Time Message Protocol)
2. HLS (HTTP Live Streaming)
3. FLV
4. Codec
5. FFMPEG
6. CDN (Content Delivery Network)

## 예상 서비스 논리 구성
실시간 스트리밍 서비스에 필요할 것으로 예상되는 개념을 정리하여 예상해본 논리적인 서비스 구성 예상도  

실제 서비스는 Web 및 Mobile UI, UI API 등도 포함이 되어 있을 것으로 생각되며  
운영상의 이유 혹은 분산 처리의 목적으로 각 서비스들이 더욱 세부적인 단위로 분리되어 있을 수 있으나  

`실시간 인터넷 방송`에 초점을 맞춰 간략하게 구성함

1. Web Service : 실시간 영상 수신 및 송출에 필요한 웹 서비스
   - Streamer로부터 RTMP 요청을 통해 영상 수신
   - 영상을 Viewer에게 HLS로 송출
   - 수신 및 송출을 담당하는 서비스는 분리하여 각자 구성할 수 있음
2. Middleware : 시스템 전반을 관리하는 서비스
   - 각 서비스간 처리 연계 Interface(Message Queue나 In-Memory DB 등을 이용할 수 있음)
   - Logging, Monitoring 서비스와 연계하여 시스템 현황 기록 및 관리(ELK 등을 이용할 수 있음)
3. Encoder : 영상 데이터 Encoding 서비스
   - RTMP로 수신 받은 영상 데이터를 HLS 응답에 이용할 `.3m8u` 및 `.ts` 파일로 Encoding
   - 영상 데이터를 VOD 서비스용 Codec으로 Encoding
   - Encoding 처리된 영상은 NFS나 Object Storage 등에 저장하여 데이터를 이용하는 서비스에서 접근하는 방식으로 처리할 수 있음
4. Broadcast Service : 방송 진행과 시청에 필요한 요청 처리
   - Streamer가 방송 정보 변경 시 Viewer의 화면에도 반영
   - 채팅 서비스
   - Viewer가 구독, 후원 시 해당 정보를 Streamer의 방송 시스템에 전달

## 궁금점
1. 스트리밍 서비스 시 많은 시청자가 방송을 시청하는 경우 발생하는 물리적인 부하는 어떻게 분산할까? (CDN? Object Storage)
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