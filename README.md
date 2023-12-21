# Live Video Streaming Service 구성해보기
- `RTMP` & `HLS` Protocol Server를 구성하여 `실시간 인터넷 방송` 서비스 구성해보기

## 개념 정리
1. RTMP (Real Time Message Protocol)
2. HLS (HTTP Live Streaming)
3. FLV
4. Codec
5. FFMPEG
6. CDN(Content Delivery Network)

## 예상 서비스 구성

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