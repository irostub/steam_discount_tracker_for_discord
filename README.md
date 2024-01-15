
# steam discount tracker for discord 
## 디스코드 웹훅을 통한 스팀 할인 목록 추적

# 기능
스팀에서 주기적으로 할인 목록을 불러와 디스코드의 웹훅으로 푸시합니다.  
스팀 할인 목록을 스냅샷 찍는 주기는 기본 30초이며 새로운 할인이 등록되었을 때만 웹훅으로 푸시됩니다.

스냅샷 주기를 변경하고싶다면 아래의 실행 옵션을 참고하여 실행 시 옵션에 추가합니다.

# 사용 방법
실행파일을 다운로드 받고 아래를 참고하여 옵션을 넣어 실행합니다.  
다운로드 링크 : [이곳을 눌러 다운받으세요](/guides/content/editing-an-existing-page)

## 실행 예시
```shell
linux_amd64_steam_discound_tracker_for_discord -webhook_url='https://discord.com/api/webhooks/00000000/00000-ABCDEFG'
```

## 실행 옵션
| 옵션  | 설명                  | 타입     | 기본값 |
|-----|---------------------|--------|----|
| webhook_url | \[필수값\] 디스코드 웹훅 URL | string |  |
| color | 디스코드 임베드 메세지 라인 색상  | int    |  15844367  |
| check_cycle | 스팀 할인 목록 갱신 주기      | int    |  30  |
| currency_symbol | 화폐 단위 표기 문자         | string |  ₩  |
