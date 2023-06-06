$(document).ready(function () {
    $(".pkDanger").click(function (e) {
        e.preventDefault();

        let id = $(this).data("id")

        let d = {
            classroomId: id,
            timestamp: Math.round(new Date().getTime() / 1000),
            sign: "",
        }
        d.sign = SignCreate(d)

        $.ajax({
            type: "POST",
            url: "/api/pk/assist",
            data: d,
            dataType: "json",
            success: function (res) {
                if (res.code == 0) {
                    toastr.success("请求成功！")
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

        return false
    });

    $("#update").click(function (e) {
        e.preventDefault();

        let url = $("#update").data("url");

        window.location.href = url
    });

    let today = $('meta[name="today"]').data("value");
    if (Number(today) == 0) {
        $("#todayNull").show();
    }else {
        $("#todayNull").hide();
    }
});