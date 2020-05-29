{{template "header.tpl"}}
<style>
td.filename {
    cursor: pointer;
}
</style>
共{{.Length}}条记录
<table class="table">
    <thead>
    <tr>
        <th>ID</th>
        <th>名称</th>
        <th>类型</th>
    </tr>
    </thead>
    <tbody>
    {{range $i, $file := .Files}}
        {{if .IsDirectory}}
            <tr>
                <td class="filename" data-href="{{$file.GetPreviewUrl}}">{{$file.Name}}</td>
                <td>-</td>
                <td>-</td>
            </tr>
        {{end}}
    {{end}}
    {{range $i, $file := .Files}}
        {{if not .IsDirectory}}
            <tr>
                <td class="filename" data-href="{{$file.GetPreviewUrl}}">{{$file.Name}}</td>
                <td>{{$file.Size}}</td>
                <td>{{$file.MimeType}}</td>
            </tr>
        {{end}}
    {{end}}
    </tbody>
</table>
<script src="https://cdn.jsdelivr.net/npm/jquery@1.12.4/dist/jquery.min.js"></script>
<script>
$(function () {
    $("table").on('click', "td.filename", function () {
        location.href = $(this).attr('data-href');
    })
});
</script>
{{template "footer.tpl"}}