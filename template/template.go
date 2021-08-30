package template

import "fmt"

var frame = `
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <title>Word文档字符数统计工具</title>
	<style>
		h2{
			text-align: center;
		}
		div{
			margin: auto;
		}
		.card{
			width: 300px;
			border-style: solid;
			border-width: 1px;
			border-color: black;
		}
		.card-head{
			text-align: center;
			border-style: solid;
			border-width: 1px;
			border-color: black;
			background: grey;
		}
		.card-body{
			height: 200px;
			text-align: center;
		}
	</style>
</head>

<body>
	</br>
	<div class="card">
		<div class="card-head">
			Word文档字符数统计工具
		</div>
		<div class="card-body">
			</br>
			<form action="/post" method="post">
				<label for="direc">请输入目录</label>
				<input type="text" name="direc" placeholder="d:\data\folder" />
				<input type="submit" value="确定" />
			</form>
			%s
		</div>
	</div>
</body>

</html>
`

// GetIndex return index html
func GetIndex() string {
	return fmt.Sprintf(frame, "")
}

var downloadAddr = `
</br>
<div>
	<a href="/download"> Result.xlsx </a>
</div>
</br>
`

// GetReply return download html
func GetReply() string {
	return fmt.Sprintf(frame, downloadAddr)
}

var alert = `
</br>
<div>
	<p style="color: red"> %s </p>
</div>
</br>
`

func GetAlert(s string) string {
	fullAlert := fmt.Sprintf(alert, s)
	return fmt.Sprintf(frame, fullAlert)
}
