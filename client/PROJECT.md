# Project 구조
tms-react-boilerplate(가칭), React, craco, tailwind(css framework), fontAwesome 적용

## Folder Tree
```bash
📦tms-react-boilerplate
 ┣ 📂.vscode
 ┣ 📂docs
 ┣ 📂node_modules                               
 ┣ 📂public                                    # React Static File
 ┃ ┣ 📜authenticate.json                       # 키클록 설정 파일 (ex, realm 명, url등)
 ┃ ┣ 📜favicon.ico
 ┃ ┣ 📜index.html
 ┃ ┗ 📜manifest.json                           # App 정보를 담은 manuifest.json 파일
 ┣ 📂src
 ┃ ┣ 📂common                                  # empty
 ┃ ┣ 📂components                              # 컴포넌트(view) 폴더
 ┃ ┃ ┣ 📂layout                                # aside / contents 영영으로 이루어진 전체 레이아웃 폴더
 ┃ ┃ ┃ ┣ 📂aside                               # aside, left 메뉴 영역
 ┃ ┃ ┃ ┃ ┣ 📜aside-left-menu-item.tsx
 ┃ ┃ ┃ ┃ ┣ 📜aside-left-menu.scss
 ┃ ┃ ┃ ┃ ┗ 📜aside-left-menu.tsx
 ┃ ┃ ┃ ┣ 📂contents                            # content, 우측 전체 화면 영역
 ┃ ┃ ┃ ┃ ┣ 📜contents-wrap.scss
 ┃ ┃ ┃ ┃ ┗ 📜contents-wrap.tsx
 ┃ ┃ ┃ ┣ 📜main-aside.tsx
 ┃ ┃ ┃ ┣ 📜main-container.tsx
 ┃ ┃ ┃ ┣ 📜main-content.tsx
 ┃ ┃ ┃ ┣ 📜main-footer.tsx
 ┃ ┃ ┃ ┣ 📜main-frame.tsx
 ┃ ┃ ┃ ┣ 📜main-wrap.tsx
 ┃ ┃ ┃ ┗ 📜under-construction.tsx
 ┃ ┃ ┣ 📂partial                               # 부분적으로, 여러군데에서 쓰일 수 있는 공통 컴포넌트
 ┃ ┃ ┃ ┗ 📜processing.tsx                      # loading(spinner) 컴포넌트
 ┃ ┃ ┗ 📂views                                 # 레이아웃 안으로 배치 될 기능 별(메뉴) 컴포넌트
 ┃ ┃ ┃ ┣ 📂a-sub-component
 ┃ ┃ ┃ ┃ ┗ 📜a-sub-component.tsx
 ┃ ┃ ┃ ┣ 📂b-sub-component
 ┃ ┃ ┃ ┃ ┗ 📜b-sub-component.tsx
 ┃ ┃ ┃ ┗ 📂main-component
 ┃ ┃ ┃ ┃ ┗ 📜main-component.tsx
 ┃ ┣ 📂service                                 # empty
 ┃ ┣ 📂store                                   # redux 관련 action, reducer 예제
 ┃ ┃ ┣ 📜ActionTest.ts
 ┃ ┃ ┣ 📜configureStore.ts
 ┃ ┃ ┣ 📜Counter.ts
 ┃ ┃ ┣ 📜index.ts
 ┃ ┃ ┗ 📜WeatherForecasts.ts
 ┃ ┣ 📜App.test.tsx
 ┃ ┣ 📜App.tsx
 ┃ ┣ 📜index.scss
 ┃ ┣ 📜index.tsx
 ┃ ┣ 📜react-app-env.d.ts
 ┃ ┣ 📜registerServiceWorker.ts                # pwa(Progressive Web Apps) 개발 관련 serviceWorker
 ┃ ┗ 📜setupProxy.js                           # proxy 설정
 ┣ 📜.editorconfig
 ┣ 📜.eslintrc.json
 ┣ 📜.gitignore
 ┣ 📜.npmrc                                    # .npmrc 파일, fontawesome pro 버전 라이센스 정보를 담고 있다.
 ┣ 📜craco.config.js                           # craco, tailwind를 사용하기 위해 설정 (postCSS 구성 재정의) 
 ┣ 📜package-lock.json
 ┣ 📜package.json
 ┣ 📜README.md
 ┣ 📜tailwind.config.js                        # css 프레임워크, tailwind 설정 파일
 ┗ 📜tsconfig.json                             # typescript 설정
```