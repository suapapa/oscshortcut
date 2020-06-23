# oscshorcut

구형 타블렛에 TouchOSC앱으로 OBS studio를 제어하기 위한 프로그램.

예로, OBS studio에 다음과 같이 단축키를 설정:

* Ctrl+Alt+9 : 방송시작
* Ctrl+Alt+0 : 방송중단

TouchOSC의 버튼 컨트롤의 주소와 값을 위의 단축키로 아래의 코드를 사용해 변환하고 있습니다.

    keymap = map[oscVal]keyComb{
        oscVal{"/1/toggle1", 0}:        keyComb{keybd_event.VK_0, true, true, false}, // 방송중단
        oscVal{"/1/toggle1", 1}:        keyComb{keybd_event.VK_9, true, true, false}, // 방송시작
    }

## reference

* [TouchOSC](https://hexler.net/products/touchosc)
