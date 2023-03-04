import axios from "axios"
import {ElMessage} from "element-plus"

axios.defaults.timeout = 60000
axios.defaults.baseURL = import.meta.env.VITE_APP_BASE_API

axios.interceptors.request.use(
    config => {
        config.headers["Content-Type"] = 'application/json;charset=UTF-8'
        return config
    },
    error => {
        return Promise.reject(error)
    }
)

axios.interceptors.response.use(response => {
        let resp = response.data
        if (resp.code !== 0) {
            ElMessage.error(resp.message)
            return
        }
        return resp
    },
    error => {
        const {r} = error
        if (r) {
            ElMessage.error(r.status)
            return Promise.reject(r.data)
        } else {
            ElMessage.error('网络请求失败')
        }
    }
)

const request = (url = '', params = {}, type = 'GET') => {
    let promise
    switch (type.toUpperCase()) {
        case 'POST': {
            promise = axios({
                method: 'POST',
                url: url,
                data: params
            })
            break
        }
        case 'PUT': {
            promise = axios({
                method: 'PUT',
                url: url,
                data: params
            })
            break
        }
        case 'PATCH': {
            promise = axios({
                method: 'PATCH',
                url: url,
                data: params
            })
            break
        }
        case 'DELETE': {
            promise = axios({
                method: 'DELETE',
                url: url,
                params: params
            })
            break
        }
        case 'GET':
        default: {
            promise = axios({
                url: url,
                params: params
            })
            break
        }
    }
    return new Promise((resolve, reject) => {
        promise.then(resp => {
            resolve(resp)
        }).catch(err => {
            reject(err)
        })
    })
}

const get = (url, params = {}) => {
    return request(url, params)
}

const post = (url, params = {}) => {
    return request(url, params, 'POST')
}

const put = (url, params = {}) => {
    return request(url, params, 'PUT')
}

const patch = (url, params = {}) => {
    return request(url, params, 'PATCH')
}

const del = (url, params = {}) => {
    return request(url, params, 'DELETE')
}

export {
    request,
    get,
    post,
    put,
    patch,
    del
}