package com.aruba.simpl.securityattributesprovider.services.impl;

import static org.junit.Assert.assertEquals;

import com.aruba.simpl.securityattributesprovider.configurations.DBSeedingProperties;
import com.aruba.simpl.securityattributesprovider.model.repositories.DBTest;
import com.aruba.simpl.securityattributesprovider.model.repositories.IdentityAttributeRepository;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.autoconfigure.validation.ValidationAutoConfiguration;
import org.springframework.boot.context.properties.EnableConfigurationProperties;
import org.springframework.context.annotation.Import;

@EnableConfigurationProperties(DBSeedingProperties.class)
@Import({SecurityAttributeProviderInitializationImpl.class, ValidationAutoConfiguration.class})
public class SecurityAttributeProviderInitializationImplEndToEndTest extends DBTest {

    @Autowired
    private IdentityAttributeRepository identityAttributeRepository;

    @Test
    public void loadContextAndInitialization() {
        assertEquals(7, identityAttributeRepository.findAll().size());
    }
}
