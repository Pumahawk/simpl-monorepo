package usersroles.controllers;

import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

import com.aruba.simpl.common.exchanges.AttachmentExchange;
import com.aruba.simpl.common.exchanges.CertificateExchange;
import com.aruba.simpl.common.exchanges.UserExchange;
import com.aruba.simpl.usersroles.configurations.MtlsClientBuilder;
import com.aruba.simpl.usersroles.service.*;
import com.aruba.simpl.usersroles.service.core.KeycloakRealmService;
import com.aruba.simpl.usersroles.service.user.KeycloakUserService;
import java.io.ByteArrayInputStream;
import java.io.FileOutputStream;
import org.apache.commons.compress.utils.IOUtils;
import org.junit.jupiter.api.Test;
import org.keycloak.admin.client.Keycloak;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.boot.autoconfigure.jdbc.DataSourceAutoConfiguration;
import org.springframework.boot.autoconfigure.ssl.SslAutoConfiguration;
import org.springframework.boot.autoconfigure.web.client.RestClientSsl;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.test.context.TestPropertySource;
import org.springframework.test.web.servlet.MockMvc;

@SpringBootTest
@AutoConfigureMockMvc
@EnableAutoConfiguration(exclude = {DataSourceAutoConfiguration.class, SslAutoConfiguration.class})
@MockBean({
    RestClientSsl.class,
    CertificateExchange.class,
    AttachmentExchange.class,
    UserExchange.class,
    CredentialService.class,
    EchoService.class,
    KeycloakInitializerService.class,
    KeycloakService.class,
    MtlsService.class,
    KeycloakRealmService.class,
    KeycloakUserService.class,
    Keycloak.class,
    MtlsClientBuilder.class,
})
@TestPropertySource(
        properties = {
            "logging.level.org.hibernate.SQL=DEBUG",
        })
public class SwaggerTest {

    Logger logger = LoggerFactory.getLogger(getClass());

    @Autowired
    MockMvc mockMvc;

    @Test
    public void generateJson() throws Exception {
        mockMvc.perform(get("/v3/api-docs")).andExpect(status().is(200)).andDo(result -> {
            var outputPath = getClass().getResource("/").getPath() + "swagger.json";
            logger.info("Output swagger path: {}", outputPath);
            IOUtils.copy(
                    new ByteArrayInputStream(result.getResponse().getContentAsByteArray()),
                    new FileOutputStream(outputPath));
        });
    }
}
