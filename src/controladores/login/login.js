

//g_requestVerificationToken
//token0_client=SMM/5SZWOp+TT9UYbRKGaY7sADI+/Nhb
//token1_client=MlmLHjhykpLol7k6nq0BgYA0plN3lIRl


function login(destnation, callback, redirectDes) {
    var name = $.trim($('#username').val());
    var psd = $('#password').val();
    var valid = validateInput(name, psd);
    if (!valid) {
        return;
    }
    if ($.isArray(g_requestVerificationToken)) {
        if (g_requestVerificationToken.length > 0) {
            if (g_password_type == '4') {
                psd = base64encode(SHA256(name + base64encode(SHA256($('#password').val())) + g_requestVerificationToken[0]));
            } else {
                psd = base64encode($('#password').val());
            }
        } else {
            setTimeout(function () {
                if (g_requestVerificationToken.length > 0) {
                    login(destnation, callback, redirectDes);
                }
            }, 50)
            return;
        }
    } else {
        psd = base64encode($('#password').val());
    }
    var request = {
        Username: name,
        Password: psd,
        password_type: g_password_type
    };
    if (valid) {
        var xmlstr = object2xml('request', request);
        log.debug('xmlstr = ' + xmlstr);
        saveAjaxData('api/user/login', xmlstr, function ($xml) {
            log.debug('api/user/login successed!');
            var ret = xml2object($xml);
            if (isAjaxReturnOK(ret)) {
                $('#username_span').text(name);
                $('#username_span').show();
                $('#logout_span').text(common_logout);
                var passwordStr = $('#password').val();
                clearDialog();
                g_main_displayingPromptStack.pop();
                startLogoutTimer(redirectDes);
                if (checkPWRemind(passwordStr)) {
                    checkDialogFlag = true;
                    showPWRemindDialog(destnation, callback);
                } else {
                    loginSwitchDoing(destnation, callback);
                }
            } else {
                if (ret.type == 'error') {
                    clearAllErrorLabel();
                    if (ret.error.code == ERROR_LOGIN_PASSWORD_WRONG) {
                        showErrorUnderTextbox('password', system_hint_wrong_password);
                        $('#password').val('');
                        $('#password').focus();
                    } else if (ret.error.code == ERROR_LOGIN_ALREADY_LOGIN) {
                        showErrorUnderTextbox('password', common_user_login_repeat);
                        $('#username').focus();
                        $('#username').val('');
                        $('#password').val('');
                    } else if (ret.error.code == ERROR_LOGIN_USERNAME_WRONG) {
                        showErrorUnderTextbox('username', settings_hint_user_name_not_exist);
                        $('#username').focus();
                        $('#username').val('');
                        $('#password').val('');
                    } else if (ret.error.code == ERROR_LOGIN_USERNAME_PWD_WRONG) {
                        showErrorUnderTextbox('password', IDS_login_username_password_wrong);
                        $('#username').focus();
                        $('#username').val('');
                        $('#password').val('');
                    } else if (ret.error.code == ERROR_LOGIN_USERNAME_PWD_ORERRUN) {
                        showErrorUnderTextbox('password', IDS_login_username_password_input_overrun);
                        $('#username').focus();
                        $('#username').val('');
                        $('#password').val('');
                    }
                }
            }
        }, {
            enc: true
        });
    }
}



function base64encode(str) {
    var out, i, len;
    var c1, c2, c3;
    const g_base64EncodeChars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/'
    len = str.length;
    i = 0;
    out = '';
    while (i < len) {
        c1 = str.charCodeAt(i++) & 0xff;
        if (i == len) {
            out += g_base64EncodeChars.charAt(c1 >> 2);
            out += g_base64EncodeChars.charAt((c1 & 0x3) << 4);
            out += '==';
            break;
        }
        c2 = str.charCodeAt(i++);
        if (i == len) {
            out += g_base64EncodeChars.charAt(c1 >> 2);
            out += g_base64EncodeChars.charAt(((c1 & 0x3) << 4) | ((c2 & 0xF0) >> 4));
            out += g_base64EncodeChars.charAt((c2 & 0xF) << 2);
            out += '=';
            break;
        }
        c3 = str.charCodeAt(i++);
        out += g_base64EncodeChars.charAt(c1 >> 2);
        out += g_base64EncodeChars.charAt(((c1 & 0x3) << 4) | ((c2 & 0xF0) >> 4));
        out += g_base64EncodeChars.charAt(((c2 & 0xF) << 2) | ((c3 & 0xC0) >> 6));
        out += g_base64EncodeChars.charAt(c3 & 0x3F);
    }
    return out;
}