# oscshorcut

구형 타블렛에 TouchOSC앱으로 OBS studio를 제어하기 위한 프로그램.

예로, OBS studio에 다음과 같이 단축키를 설정:

* Ctrl+Alt+9 : 방송시작
* Ctrl+Alt+0 : 방송중단

TouchOSC의 버튼 컨트롤의 주소와 값을 위의 단축키로 매핑하기위해 `shortcuts.json`을
아래와 같이 정의해 사용:

    [
        {
            "oscVal": {
                "addr": "/1/multipush2/1/1",
                "val": 4
            },
            "keyComb": {
                "description": "sceen 1",
                "key": "1",
                "ctrl": true,
                "alt": true,
                "shift": false
            }
        },
        {
            "oscVal": {
                "addr": "/1/multipush2/2/1",
                "val": 4
            },
            "keyComb": {
                "description": "sceen 2",
                "key": "2",
                "ctrl": true,
                "alt": true,
                "shift": false
            }
        },
        ...
    ]


## reference

* [TouchOSC](https://hexler.net/products/touchosc)
