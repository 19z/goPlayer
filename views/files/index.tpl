{{template "header.tpl"}}
<style>
    td.filename {
        cursor: pointer;
    }
</style>

<div class="container">
    <div class="row">
        <div class="col-md-10">
            <div class="panel panel-default">
                <div class="panel-heading">共{{.Length}}条记录</div>
                <table class="table">
                    <thead>
                    <tr>
                        <th>文件名</th>
                        <th>大小</th>
                        <th>更新时间</th>
                    </tr>
                    </thead>
                    <tbody>
                    {{if .Parent }}
                        <tr>
                            <td class="filename" data-href="{{.Parent.GetPreviewUrl}}">../ 上一层目录</td>
                            <td>-</td>
                            <td>{{.Parent.ModTime|TimeFormat}}</td>
                        </tr>
                    {{end}}
                    {{range $i, $file := .Files}}
                        {{if .IsDirectory}}
                            <tr>
                                <td class="filename" data-href="{{$file.GetPreviewUrl}}">{{$file.Name}}</td>
                                <td>-</td>
                                <td>{{$file.ModTime|TimeFormat}}</td>
                            </tr>
                        {{end}}
                    {{end}}
                    {{range $i, $file := .Files}}
                        {{if not .IsDirectory}}
                            <tr>
                                <td class="filename" data-href="{{$file.GetPreviewUrl}}">
                                    {{if $file.IsVideo}}
                                        <img width="30px" src="{{$file.GetPreviewUrl}}?mode=pic" alt="">
                                    {{end}}
                                    {{$file.Name}}
                                </td>
                                <td>{{$file.Size|FileSizeFormat}}</td>
                                <td>{{$file.ModTime|TimeFormat}}</td>
                            </tr>
                        {{end}}
                    {{end}}
                    </tbody>
                </table>
            </div>
        </div>
        <div class="col-md-2">
            <div class="panel panel-default">
                <div class="panel-heading">上传</div>
                <div class="panel-body">
                    <form method="post">
                        <div class="form-group">
                            <label>名称</label>
                            <input name="name" type="text" class="form-control">
                        </div>
                        <div class="form-group">
                            <label>类型</label>
                            <select name="type" class="form-control">
                                <option value="local">本地</option>
                            </select>
                        </div>
                        <div class="form-group">
                            <label>地址</label>
                            <input name="path" type="text" class="form-control">
                        </div>
                        <input type="submit" class="btn btn-default" value="保存">
                    </form>
                </div>
            </div>
        </div>
    </div>

</div>
<script src="https://cdn.jsdelivr.net/npm/jquery@1.12.4/dist/jquery.min.js"></script>
<script>
    $(function () {
        $("table").on('click', "td.filename", function () {
            location.href = $(this).attr('data-href');
        })
    });
</script>
{{template "footer.tpl"}}