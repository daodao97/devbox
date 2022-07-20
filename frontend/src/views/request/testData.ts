import type { Group } from "./model"

const preScript = `
export default {
    preRequestHandle(options) {
        return options
    },
    responseHandle(body) {
        return {"a":1}
    },
}   
`

export const data: Group[] = [
    {
        id: 2,
        name: 'Group1',
        children: [
            {
                method: "GET",
                name: '/ok/test1',
                url: 'https://api.github.com/repos/rogchap/wombat/releases/latest',
                script: preScript
            },
            {
                method: "POST",
                name: '/ok/test2',
                url: 'https://api.github.com/repos/rogchap/wombat/releases/latest'
            },
            {
                method: "DELETE",
                name: '/ok/test3',
                url: 'https://api.github.com/repos/rogchap/wombat/releases/latest'
            },
        ],
    },
]