package properties

import com.natpryce.konfig.getValue
import com.natpryce.konfig.stringType


object System {
    val environment by stringType
    val api_version by stringType
}