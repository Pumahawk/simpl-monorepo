package com.aruba.simpl.securityattributesprovider.services.impl;

import static org.junit.Assert.assertEquals;
import static org.mockito.Mockito.*;

import com.aruba.simpl.common.exchanges.onboarding.ParticipantTypeExchange;
import com.aruba.simpl.common.model.dto.ParticipantTypeDTO;
import com.aruba.simpl.securityattributesprovider.configurations.DBSeedingProperties;
import com.aruba.simpl.securityattributesprovider.model.mappers.IdentityAttributeMapperImpl;
import com.aruba.simpl.securityattributesprovider.model.repositories.DBTest;
import com.aruba.simpl.securityattributesprovider.model.repositories.IdentityAttributeRepository;
import java.util.Set;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.autoconfigure.validation.ValidationAutoConfiguration;
import org.springframework.boot.context.properties.EnableConfigurationProperties;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.context.ApplicationContext;
import org.springframework.context.annotation.AnnotationConfigApplicationContext;
import org.springframework.context.annotation.Import;

@EnableConfigurationProperties(DBSeedingProperties.class)
@Import({ValidationAutoConfiguration.class, IdentityAttributeMapperImpl.class})
public class SecurityAttributeProviderInitializationImplEndToEndTest extends DBTest {

    @Autowired
    private ApplicationContext applicationContext;

    @MockBean
    private ParticipantTypeExchange participantTypeExchange;

    @Autowired
    private IdentityAttributeRepository identityAttributeRepository;

    @Test
    public void loadContextAndInitialization() {
        var validParticipantTypes = getValidParticipantTypes();
        when(participantTypeExchange.getParticipantTypes()).thenReturn(validParticipantTypes);
        initializationSecurityAttributes();
        assertEquals(7, identityAttributeRepository.findAll().size());
    }

    private Set<ParticipantTypeDTO> getValidParticipantTypes() {
        return Set.of(
                createParticipantTypeDTO("APPLICATION_PROVIDER"),
                createParticipantTypeDTO("CONSUMER"),
                createParticipantTypeDTO("DATA_PROVIDER"),
                createParticipantTypeDTO("INFRASTRUCTURE_PROVIDER"));
    }

    private ParticipantTypeDTO createParticipantTypeDTO(String value) {
        var pt = new ParticipantTypeDTO();
        pt.setValue(value);
        return pt;
    }

    private void initializationSecurityAttributes() {
        try (var testContext = new AnnotationConfigApplicationContext()) {
            testContext.setParent(applicationContext);
            testContext.register(SecurityAttributeProviderInitializationImpl.class);
            testContext.refresh();
        }
    }
}
