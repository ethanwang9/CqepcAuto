$(document).ready(function () {
    $("form").submit(function (e) {
        e.preventDefault();

        let d = {}
        $("form").serializeArray().forEach((v) => {
            d[v.name] = v.value
        });
        d["timestamp"] = Math.round(new Date().getTime() / 1000)
        d["sign"] = SignCreate(d)

        $.ajax({
            type: "POST",
            url: "/install/verify",
            data: d,
            dataType: "json",
            beforeSend: function () {
                $("#loadingMsg").modal({
                    backdrop: 'static',
                    keyboard: false,
                    show: true,
                });
            },
            success: function (res) {
                if (res.code == 0) {
                    toastr.success("验证成功，3秒后进入后台管理界面！")
                    setTimeout(() => {
                        window.location.href = "/admin/index?t=" + Math.round(new Date().getTime() / 1000)
                    }, 3000);
                } else {
                    toastr.warning(res.msg)
                }
            },
            error: function (res) {
                toastr.error("网络请求失败，请检测程序是否在运行！")
            },
            complete: function () {
                setTimeout(() => {
                    $("#loadingMsg").modal('hide')
                }, 500);
            },
        });

        return false
    });
});