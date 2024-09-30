package com.aruba.simpl.usersroles.controllers;

import com.aruba.simpl.backend.common.SwaggerGeneratorTest;
import com.aruba.simpl.common.exchanges.AttachmentExchange;
import com.aruba.simpl.common.exchanges.CertificateExchange;
import com.aruba.simpl.common.exchanges.UserExchange;
import com.aruba.simpl.usersroles.configurations.MtlsClientBuilder;
import com.aruba.simpl.usersroles.configurations.RoleInitializer;
import com.aruba.simpl.usersroles.service.*;
import com.aruba.simpl.usersroles.service.core.KeycloakRealmService;
import com.aruba.simpl.usersroles.service.user.KeycloakUserService;
import org.keycloak.admin.client.Keycloak;
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
    IdentityAttributeService.class,
    RoleInitializer.class,
    RoleService.class,
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
public class UsersRolesSwaggerTest extends SwaggerGeneratorTest {}
