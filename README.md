
# steam discount tracker for discord 
## 디스코드 웹훅을 통한 스팀 할인 목록 추적
스팀에서 주기적으로 할인 목록을 불러와 디스코드의 웹훅으로 푸시합니다.  
스팀 할인 목록을 스냅샷 찍는 주기는 기본 30초이며 새로운 할인이 등록되었을 때만 웹훅으로 푸시됩니다.

스냅샷 주기를 변경하고싶다면 아래의 실행 옵션을 참고하여 실행 시 옵션에 추가합니다.
![image](https://github.com/irostub/steam_discount_tracker_for_discord/assets/61470181/3265d622-30f7-4bd1-827c-f32d9b63cfcd)

# 실행 파일 사용 방법
실행파일을 다운로드 받고 아래를 참고하여 옵션을 넣어 실행합니다.  
운영체제와 아키텍처에 맞는 tar 파일을 아래 링크로가서 다운로드 받습니다.

예)   
linux 운영체제, amd cpu :linux_amd64.tar.gz  
linux 운영체제, arm cpu :linux_arm64.tar.gz

다운로드 링크 : [https://github.com/irostub/steam_discount_tracker_for_discord/releases](https://github.com/irostub/steam_discount_tracker_for_discord/releases)



## 실행 예시
실행 옵션에 -webhook_url='디스코드 웹훅 URL' 을 포함하여 실행해야합니다.
```shell
wget https://github.com/irostub/steam_discount_tracker_for_discord/releases/download/1.3/steam_discount_tracker_for_discord_1.3_linux_amd64.tar.gz

tar -xzf steam_discount_tracker_for_discord_1.3_linux_amd64.tar.gz

./steam_discount_tracker_for_discord_1.3_linux_amd64 -webhook_url='https://discord.com/api/webhooks/00000000/00000-ABCDEFG'
```

## 실행 옵션
| 옵션  | 설명                  | 타입     | 기본값 |
|-----|---------------------|--------|----|
| webhook_url | \[필수값\] 디스코드 웹훅 URL | string |  |
| color | 디스코드 임베드 메세지 라인 색상  | int    |  15844367  |
| check_cycle | 스팀 할인 목록 갱신 주기      | int    |  30  |
| currency_symbol | 화폐 단위 표기 문자         | string |  ₩  |


---
# Docker 통한 사용 방법
Docker Image 는 dockerhub 에 public 으로 공개되어있습니다.  

이미지 명은 아래와 같습니다.

이미지 명 : irostub/steam_discount_tracker_for_discord

## 실행 예시
```shell
docker run -d irostub/steam_discount_tracker_for_discord -e WEBHOOK_URL='https://discord.com/api/webhooks/00000000/00000-ABCDEFG'
```

## Docker ENV 옵션 목록
| 옵션  | 설명                  | 타입     | 기본값 |
|-----|---------------------|--------|----|
| WEBHOOK_URL | \[필수값\] 디스코드 웹훅 URL | string |  |
| COLOR | 디스코드 임베드 메세지 라인 색상  | int    |  15844367  |
| CHECK_CYCLE | 스팀 할인 목록 갱신 주기      | int    |  30  |
| CURRENCY_SYMBOL | 화폐 단위 표기 문자         | string |  ₩  |
