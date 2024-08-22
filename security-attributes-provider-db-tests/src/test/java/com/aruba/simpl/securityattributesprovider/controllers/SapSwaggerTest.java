package com.aruba.simpl.securityattributesprovider.controllers;

import com.aruba.simpl.backend.common.SwaggerGeneratorTest;
import com.aruba.simpl.common.exchanges.AttachmentExchange;
import com.aruba.simpl.common.exchanges.CertificateExchange;
import com.aruba.simpl.common.exchanges.UserExchange;
import com.aruba.simpl.securityattributesprovider.services.CliService;
import com.aruba.simpl.securityattributesprovider.services.IdentityAttributeService;
import com.aruba.simpl.securityattributesprovider.services.ParticipantService;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.boot.autoconfigure.jdbc.DataSourceAutoConfiguration;
import org.springframework.boot.autoconfigure.ssl.SslAutoConfiguration;
import org.springframework.boot.autoconfigure.web.client.RestClientSsl;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.test.context.TestPropertySource;

@SpringBootTest
@AutoConfigureMockMvc
@EnableAutoConfiguration(exclude = {DataSourceAutoConfiguration.class, SslAutoConfiguration.class})
@MockBean({
    RestClientSsl.class,
    CertificateExchange.class,
    AttachmentExchange.class,
    UserExchange.class,
    CliService.class,
    IdentityAttributeService.class,
    ParticipantService.class,
})
@TestPropertySource(
        properties = {
            "logging.level.org.hibernate.SQL=DEBUG",
        })
public class SapSwaggerTest extends SwaggerGeneratorTest {}
