package com.aruba.simpl.securityattributesprovider.controllers;

import com.aruba.simpl.common.exchanges.AttachmentExchange;
import com.aruba.simpl.common.exchanges.CertificateExchange;
import com.aruba.simpl.common.exchanges.UserExchange;
import org.junit.jupiter.api.AfterAll;
import org.junit.jupiter.api.BeforeAll;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.boot.autoconfigure.ssl.SslAutoConfiguration;
import org.springframework.boot.autoconfigure.web.client.RestClientSsl;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.boot.test.mock.mockito.MockBeans;
import org.springframework.test.annotation.DirtiesContext;
import org.springframework.test.context.DynamicPropertyRegistry;
import org.springframework.test.context.DynamicPropertySource;
import org.springframework.test.context.TestPropertySource;
import org.testcontainers.containers.PostgreSQLContainer;

@SpringBootTest
@AutoConfigureMockMvc
@EnableAutoConfiguration(exclude = {SslAutoConfiguration.class})
@MockBeans({
    @MockBean(RestClientSsl.class),
    @MockBean(CertificateExchange.class),
    @MockBean(AttachmentExchange.class),
    @MockBean(UserExchange.class)
})
@TestPropertySource(
        properties = {
            "logging.level.org.hibernate.SQL=DEBUG",
        })
@DirtiesContext
public class EndToEndTest {

    static PostgreSQLContainer<?> postgres = new PostgreSQLContainer<>("postgres:16-alpine");

    @BeforeAll
    static void beforeAll() {
        postgres.start();
    }

    @AfterAll
    static void afterAll() {
        postgres.stop();
    }

    @DynamicPropertySource
    static void configureProperties(DynamicPropertyRegistry registry) {
        registry.add("spring.datasource.url", postgres::getJdbcUrl);
        registry.add("spring.datasource.username", postgres::getUsername);
        registry.add("spring.datasource.password", postgres::getPassword);
        registry.add("spring.datasource.driver-class-name", postgres::getDriverClassName);
    }
}
