import {describe, expect, it,} from 'vitest'

import {EchoResponse, EchoService} from './api/echo/v1/echo.pb'
import {InitReq} from "./api/fetch.pb";


describe(`invoke echo`, () => {
    const init: InitReq = {
        pathPrefix: 'http://localhost:8081',
    }

    it(`with fetch`, async () => {
        const reply = await fetch('http://localhost:8081/v1/example/echo', {
            method: 'POST',
            body: JSON.stringify({
                message: 'hello'
            })
        })
        const data = await reply.json()
        expect(data.message).toEqual('hello')
    })

    it('with generate client', async () => {

        const resp = await EchoService.Echo({
            message: 'hello',
        }, init)

        expect(resp.message).toEqual('hello')
    })

    it(`invoke with streaming`, async () => {
        let count = 0

        function onMessage(message: EchoResponse) {
            count++
            // expect(message.message).toEqual('hello')
            console.log(1)
        }

        await EchoService.EchoStreaming({
            message: 'hello',
        }, onMessage, init)
    })
})