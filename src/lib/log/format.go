package log

const (
	loggerConfig = `
<seelog type="asyncloop">
	<outputs formatid="main">
		<console/>
		<filter levels="trace,debug">
			<rollingfile type="size" filename="log/init-server-debug.log" maxsize="104857600" maxrolls="10"/>
		</filter>
		<filter levels="info, warn, error, critical">
			<rollingfile type="date" filename="log/init-server.log" datepattern="2006-01-02" maxrolls="10"/>
			<file path="log/init-server.log"/>
		</filter>
	</outputs>
	<formats>
	<format id="main" format="[%Date %Time] [%LEVEL] - %File[%Line] %Func - %Msg%n"/>
	</formats>
</seelog>
`
)
