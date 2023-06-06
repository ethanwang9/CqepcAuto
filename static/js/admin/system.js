$(document).ready(function () {
    let loginType = $("#SysLoginType").data("value");
    console.log(loginType)
    $("#SysLoginType").find(`option[value="${loginType}"]`).attr("selected", true);

    $("#getCode").on('click', function (e) {
        e.stopPropagation();
        e.preventDefault();

        let d = {
            token: $("#DingTalkWebhook").val(),
            secret: $("#DingTalkSecret").val(),
            phone: $("#DingTalkPhone").val(),
            sign: "",
            timestamp: Math.round(new Date().getTime() / 1000),
        }

        d.sign = SignCreate(d)

        $.ajax({
            type: "POST",
            url: "/api/code/update",
            data: d,
            dataType: "json",
            success: function (res) {
                if (res.code == 0) {
                    toastr.success("获取验证码成功, 请注意查看钉钉消息！")
                } else {
                    toastr.warning(res.msg)
                }
            },
            error: function (res) {
                toastr.error("网络请求失败，请检测程序是否在运行！")
            },
        });
        return
    });

    $("form").submit(function (e) {
        e.preventDefault();
        e.stopPropagation();

        let d = {}
        $("form").serializeArray().forEach((v) => {
            d[v.name] = v.value
        });
        if(d["SysLoginType"] == "account" && d["StudentOpenid"].length == 0) {
            d["StudentOpenid"] = "NULL"
        }
        
        d["timestamp"] = Math.round(new Date().getTime() / 1000)
        d["sign"] = SignCreate(d)

        $.ajax({
            type: "POST",
            url: "/admin/system",
            data: d,
            dataType: "json",
            success: function (res) {
                if (res.code == 0) {
                    toastr.success("保存成功！")
                    setTimeout(() => {
                        window.location.reload()
                    }, 2000);
                } else {
                    toastr.warning(res.msg)
                }
            },
            error: function (res) {
                toastr.error("网络请求失败，请检测程序是否在运行！")
            },
        });

        return
    });
});