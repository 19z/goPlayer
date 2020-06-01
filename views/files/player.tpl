<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>{{.stat.Name}}</title>
    <script src="https://cdn.jsdelivr.net/npm/hls.js@0.13.2/dist/hls.min.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/dplayer@1.25.1/dist/DPlayer.min.js"></script>
    <style>
        body, html {
            margin: 0;
        }
    </style>
</head>
<body>
<div id="dplayer"></div>
<script>
    //http://dplayer.js.org/zh/guide.html
    var Splat = "{{.Splat}}";
    const dp = new DPlayer({
        container: document.getElementById('dplayer'),
        autoplay: true,
        lang: 'zh-cn',
        playbackSpeed: [0.5, 0.75, 1, 1.2, 1.4, 1.6, 1.7, 1.8, 2, 2.5, 3],
        video: {
            quality: [
                {
                    name: '原画',
                    url: "/files/" + Splat,
                    type: 'auto',
                },
                {
                    name: '1280P',
                    url: "/video/" + Splat + "?mode=m3u8&w=1280",
                    type: 'hls',
                },
                {
                    name: '720P',
                    url: "/video/" + Splat + "?mode=m3u8&w=720",
                    type: 'hls',
                },
                {
                    name: '640P',
                    url: "/video/" + Splat + "?mode=m3u8&w=640",
                    type: 'hls',
                },
                {
                    name: '320P',
                    url: "/video/" + Splat + "?mode=m3u8&w=320",
                    type: 'hls',
                },
            ],
            defaultQuality: 0,
            pic: "/video/" + Splat + "?mode=pic",
        },
    });
</script>
</body>
</html>