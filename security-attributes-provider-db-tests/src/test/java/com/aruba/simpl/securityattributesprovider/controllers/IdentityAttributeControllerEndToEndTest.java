package com.aruba.simpl.securityattributesprovider.controllers;

import static com.aruba.simpl.common.test.TestUtil.a;
import static org.assertj.core.api.Assertions.assertThat;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.put;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

import com.aruba.simpl.common.model.enums.ParticipantType;
import com.aruba.simpl.common.test.WithMockSecurityContext;
import com.aruba.simpl.securityattributesprovider.model.entities.IdentityAttribute;
import com.aruba.simpl.securityattributesprovider.model.repositories.IdentityAttributeRepository;
import com.aruba.simpl.securityattributesprovider.utils.TransactionalUtils;
import com.fasterxml.jackson.databind.ObjectMapper;
import jakarta.persistence.EntityManager;
import java.util.Arrays;
import java.util.LinkedList;
import java.util.UUID;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;

public class IdentityAttributeControllerEndToEndTest extends EndToEndTest {

    @Autowired
    MockMvc mockMvc;

    @Autowired
    private IdentityAttributeRepository repository;

    @Autowired
    private EntityManager manager;

    @Autowired
    private TransactionalUtils tr;

    @Test
    @WithMockSecurityContext(roles = "IATTR_M")
    public void updateAssignableParameter() throws Exception {

        var ids = new LinkedList<String>();

        tr.transactional(() -> {
            var val1 = a(IdentityAttribute.class);
            var val2 = a(IdentityAttribute.class);
            var val3 = a(IdentityAttribute.class);

            val1.setAssignableToRoles(false);
            val1.setId(null);
            val2.setAssignableToRoles(false);
            val2.setId(null);
            val3.setAssignableToRoles(false);
            val3.setId(null);

            repository.save(val1);
            repository.save(val2);
            repository.save(val3);

            manager.flush();
            manager.clear();

            ids.add(val1.getId().toString());
            ids.add(val2.getId().toString());
            ids.add(val3.getId().toString());
        });

        var idsUpdate = ids.subList(0, 2);

        mockMvc.perform(put("/identity-attribute/assignable/true")
                        .header("Content-type", MediaType.APPLICATION_JSON_VALUE)
                        .content(new ObjectMapper().writeValueAsString(idsUpdate)))
                .andExpect(status().is(200));

        tr.transactional(() -> {
            var val1 = repository.getReferenceById(UUID.fromString(ids.get(0)));
            var val2 = repository.getReferenceById(UUID.fromString(ids.get(1)));
            var val3 = repository.getReferenceById(UUID.fromString(ids.get(2)));

            assertThat(val1.isAssignableToRoles()).isTrue();
            assertThat(val2.isAssignableToRoles()).isTrue();
            assertThat(val3.isAssignableToRoles()).isFalse();
        });
    }

    @Test
    @WithMockSecurityContext(roles = "IATTR_M")
    public void addParticipantType() throws Exception {

        var ids = new LinkedList<String>();

        tr.transactional(() -> {
            var val1 = a(IdentityAttribute.class);
            var val2 = a(IdentityAttribute.class);
            var val3 = a(IdentityAttribute.class);
            var val4 = a(IdentityAttribute.class);

            val1.setId(null);
            val1.getParticipantTypes().clear();
            val1.getParticipantTypes().add(ParticipantType.CONSUMER);
            val2.setId(null);
            val2.getParticipantTypes().clear();
            val2.getParticipantTypes().add(ParticipantType.APPLICATION_PROVIDER);
            val3.setId(null);
            val3.getParticipantTypes().clear();
            val4.setId(null);
            val4.getParticipantTypes().clear();

            repository.save(val1);
            repository.save(val2);
            repository.save(val3);
            repository.save(val4);

            manager.flush();
            manager.clear();

            ids.add(val1.getId().toString());
            ids.add(val2.getId().toString());
            ids.add(val3.getId().toString());
            ids.add(val4.getId().toString());
        });

        var idsUpdate = ids.subList(0, 3);

        mockMvc.perform(put("/identity-attribute/addParticipantType/" + ParticipantType.DATA_PROVIDER)
                        .header("Content-type", MediaType.APPLICATION_JSON_VALUE)
                        .content(new ObjectMapper().writeValueAsString(idsUpdate)))
                .andExpect(status().is(200));

        tr.transactional(() -> {
            var val1 = repository.getReferenceById(UUID.fromString(ids.get(0)));
            var val2 = repository.getReferenceById(UUID.fromString(ids.get(1)));
            var val3 = repository.getReferenceById(UUID.fromString(ids.get(2)));
            var val4 = repository.getReferenceById(UUID.fromString(ids.get(3)));

            assertThat(val1.getParticipantTypes()
                            .containsAll(Arrays.asList(ParticipantType.CONSUMER, ParticipantType.DATA_PROVIDER)))
                    .isTrue();
            assertThat(val2.getParticipantTypes()
                            .containsAll(
                                    Arrays.asList(ParticipantType.APPLICATION_PROVIDER, ParticipantType.DATA_PROVIDER)))
                    .isTrue();
            assertThat(val3.getParticipantTypes().contains(ParticipantType.DATA_PROVIDER))
                    .isTrue();
            assertThat(val4.getParticipantTypes().isEmpty()).isTrue();
        });
    }
}
