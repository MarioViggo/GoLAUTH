import { validPassword, validEmail } from "../regex"

export const validaEmail = (email) => {
    if (!validEmail.test(email)) {
        return false
    }
    return true
}

export const validaPassword = (password) => {
    if (!validPassword.test(password)) {
        return false
    }
    return true
}