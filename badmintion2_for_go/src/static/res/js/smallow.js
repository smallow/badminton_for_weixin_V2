/**
 * Created by wanghuidong on 2017/2/24.
 */

/**
 * 消息提示 type "warn" "success"
 * @param msg
 * @param type
 */
var $loadingToast = $('#loadingToast');
var $toast = $('#toast');

function toastMsg(msg, type) {
    if ($toast.css('display') != 'none') return;
    if (type == "warn") {
        $("#toast_msg_icon").attr("class", "weui-icon-warn weui-icon_msg weui-icon_toast");
    } else if (type == "success") {
        $("#toast_msg_icon").attr("class", "weui-icon-success-no-circle weui-icon_toast");
    }

    $("#toast_msg_icon").css("font-size", "50px");
    $("#toast_msg").html(msg);
    $toast.fadeIn(100);
    setTimeout(function () {
        $toast.fadeOut(100);
    }, 2000);
}

function ajaxLoading() {
    if ($loadingToast.css('display') != 'none') return;
    $loadingToast.fadeIn(100);
}
function ajaxEnd() {
    $loadingToast.fadeOut(100);
}
