package com.milesjpool.kahd.api

import com.natpryce.konfig.ConfigurationProperties
import com.natpryce.konfig.getValue
import com.natpryce.konfig.stringType
import io.ktor.application.call
import io.ktor.http.ContentType
import io.ktor.response.respondText
import io.ktor.routing.get
import io.ktor.routing.routing
import io.ktor.server.engine.embeddedServer
import io.ktor.server.netty.Netty

val version by stringType

fun main(args: Array<String>) {

    val buildProperties = ConfigurationProperties.fromResource("build.properties")

    val server = embeddedServer(Netty, 8080) {
        routing {
            get("/") {
                call.respondText("Hello \uD83D\uDE48!", ContentType.Text.Html)
            }
            get("/version") {
                call.respondText(buildProperties[version], ContentType.Text.Html)
            }

        }
    }
    server.start(wait = true)
}