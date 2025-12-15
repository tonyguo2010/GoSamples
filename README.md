# Go Sample Projects
**1. Cobra**
   - go run main.go first first
   - go run main.go first second
   - to prepared dependancies:
     - go mod init Cobra
     - go mod tidy

**2. Aqi**
   - ![img](./02.aqi/aqi.png)
   - to prepared dependancies:
     - go mod init Aqi
     - go get -u github.com/wonli/aqi
   - go run main.go
   - to test this sample, use use [wscat](https://github.com/websockets/wscat)
      - wscat -c ws://localhost:2015/ws
        - {"action":"hi"}
          
**3. JSON**
   - go run .
