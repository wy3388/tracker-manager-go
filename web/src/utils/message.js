import {ElMessage} from "element-plus";

const show = (resp, callback = null) => {
    if (resp.code === 0) {
        ElMessage.success({
            message: resp.message,
            duration: 1000,
            onClose: () => {
                if (callback instanceof Function && typeof callback != 'undefined') {
                    callback()
                }
            }
        })
    } else {
        ElMessage.error(resp.message)
    }
}

export default {
    show
}