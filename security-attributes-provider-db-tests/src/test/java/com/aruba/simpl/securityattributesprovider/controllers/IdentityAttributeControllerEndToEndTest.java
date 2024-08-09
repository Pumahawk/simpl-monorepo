package com.aruba.simpl.securityattributesprovider.controllers;

import static com.aruba.simpl.common.test.TestUtil.a;
import static org.assertj.core.api.Assertions.assertThat;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.put;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

import com.aruba.simpl.common.model.enums.ParticipantType;
import com.aruba.simpl.common.test.WithMockSecurityContext;
import com.aruba.simpl.securityattributesprovider.model.entities.IdentityAttribute;
import com.aruba.simpl.securityattributesprovider.model.repositories.IdentityAttributeRepository;
import com.aruba.simpl.securityattributesprovider.utils.TransactionalUtils;
import com.fasterxml.jackson.databind.ObjectMapper;
import jakarta.persistence.EntityManager;
import java.util.*;
import java.util.function.Consumer;
import org.junit.jupiter.api.BeforeEach;
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

    @BeforeEach
    public void cleanDB() {
        tr.transactional(() -> {
            repository.deleteAll();
            manager.flush();
            manager.clear();
        });
    }

    @Test
    @WithMockSecurityContext(roles = "IATTR_M")
    public void updateAssignableParameter() throws Exception {

        var iau = new IAUtil();
        var ids = iau.getIds();

        tr.transactional(() -> {
            iau.createIA("val1", ia -> ia.setAssignableToRoles(false));
            iau.createIA("val2", ia -> ia.setAssignableToRoles(false));
            iau.createIA("val3", ia -> ia.setAssignableToRoles(false));
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

        var iau = new IAUtil();
        var ids = iau.getIds();

        tr.transactional(() -> {
            iau.createIA("val1", ParticipantType.CONSUMER);
            iau.createIA("val2", ParticipantType.APPLICATION_PROVIDER);
            iau.createIA("val3");
            iau.createIA("val4");
        });

        var idsUpdate = ids.subList(0, 3);

        mockMvc.perform(put("/identity-attribute/addParticipantType/" + ParticipantType.DATA_PROVIDER)
                        .header("Content-type", MediaType.APPLICATION_JSON_VALUE)
                        .content(new ObjectMapper().writeValueAsString(idsUpdate)))
                .andExpect(status().is(200));

        tr.transactional(() -> {
            iau.loadEntities();
            var val1 = iau.get(0);
            var val2 = iau.get(1);
            var val3 = iau.get(2);
            var val4 = iau.get(3);

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

    @Test
    @WithMockSecurityContext(roles = "IATTR_M")
    public void search() throws Exception {

        var iau = new IAUtil();
        var ids = iau.getIds();

        tr.transactional(() -> {
            iau.createIA("val1", ParticipantType.DATA_PROVIDER);
            iau.createIA("val2", ParticipantType.CONSUMER, ParticipantType.INFRASTRUCTURE_PROVIDER);
            iau.createIA("val3", ParticipantType.APPLICATION_PROVIDER);
            iau.createIA("val4", ParticipantType.APPLICATION_PROVIDER, ParticipantType.CONSUMER);
            iau.createIA("val5", ParticipantType.INFRASTRUCTURE_PROVIDER);
            iau.createIA("val6");
        });

        var resp = mockMvc.perform(get("/identity-attribute/search?sort=code,asc&participantTypeNotIn="
                        + ParticipantType.APPLICATION_PROVIDER))
                .andExpect(status().is(200));

        var responseNode =
                new ObjectMapper().readTree(resp.andReturn().getResponse().getContentAsString());

        var content = responseNode.withArrayProperty("content");
        assertThat(responseNode.at("/page/totalElements").asInt()).isEqualTo(4);
        assertThat(content.get(0).get("id").textValue()).isEqualTo(ids.get(0));
        assertThat(content.get(1).get("id").textValue()).isEqualTo(ids.get(1));
        assertThat(content.get(2).get("id").textValue()).isEqualTo(ids.get(4));
        assertThat(content.get(3).get("id").textValue()).isEqualTo(ids.get(5));
    }

    private class IAUtil {

        private final List<String> ids = new LinkedList<>();
        private List<IdentityAttribute> entities = Collections.emptyList();

        public List<String> getIds() {
            return ids;
        }

        public IdentityAttribute get(int i) {
            return entities.stream()
                    .filter(ia -> ia.getId().toString().equals(ids.get(i)))
                    .findAny()
                    .orElseThrow();
        }

        public List<IdentityAttribute> loadEntities() {
            this.entities =
                    repository.findAllById(ids.stream().map(UUID::fromString).toList());
            return getEntities();
        }

        public List<IdentityAttribute> getEntities() {
            return Collections.unmodifiableList(entities);
        }

        private IdentityAttribute createIA(String code, ParticipantType... pt) {
            return createIA(code, ia -> {}, pt);
        }

        private IdentityAttribute createIA(String code, Consumer<IdentityAttribute> consumer, ParticipantType... pt) {
            var ia = a(IdentityAttribute.class);
            ia.setId(null);
            ia.setCode(code);
            ia.getParticipantTypes().clear();
            ia.getParticipantTypes().addAll(Arrays.asList(pt));

            consumer.accept(ia);

            repository.save(ia);

            manager.flush();
            manager.clear();

            ids.add(ia.getId().toString());

            return ia;
        }
    }
}
