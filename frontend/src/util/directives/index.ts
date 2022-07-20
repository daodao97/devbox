import { Directive, App } from 'vue'

const vFocus: Directive = {
    mounted(el: HTMLElement) {
        if (['textarea', 'input'].includes(el.tagName)) {
            el.focus()
        }
        const textarea = el.getElementsByTagName('textarea')
        if (textarea.length > 0) {
            textarea[0].focus()
        }
        const input = el.getElementsByTagName('input')
        if (input.length > 0) {
            input[0].focus()
        }
    }
}

export default {
    install(app: App) {
        app.directive('focus', vFocus)
    }
}