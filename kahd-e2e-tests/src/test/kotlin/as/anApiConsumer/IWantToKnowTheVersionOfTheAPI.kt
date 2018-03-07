package `as`.anApiConsumer

import com.natpryce.konfig.ConfigurationProperties
import com.natpryce.konfig.ConfigurationProperties.Companion.systemProperties
import io.restassured.RestAssured.`when`
import org.hamcrest.Matchers.equalTo
import org.junit.Test
import properties.Environment
import properties.System
import java.net.URI

class `IWantToKnowTheVersionOfTheAPI` {

    private val env = systemProperties()[System.environment]
    private val apiVersion = systemProperties()[System.api_version]
    private val envProperties = ConfigurationProperties.fromResource("$env.properties")

    private val apiHost = envProperties[Environment.api_host]
    private val apiUrl = URI.create("http://$apiHost")

    @Test
    fun `It has a version endpoint`() {
        val versionUrl = apiUrl.resolve("/version")
        `when`().
            get(versionUrl).
        then().
            statusCode(200).
            content(equalTo(apiVersion))
    }
}