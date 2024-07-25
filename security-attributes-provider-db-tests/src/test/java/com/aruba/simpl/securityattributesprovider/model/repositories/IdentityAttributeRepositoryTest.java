package com.aruba.simpl.securityattributesprovider.model.repositories;

import static com.aruba.simpl.common.test.TestUtil.a;
import static org.assertj.core.api.Assertions.assertThat;

import com.aruba.simpl.securityattributesprovider.model.entities.IdentityAttribute;
import jakarta.persistence.EntityManager;
import java.util.Arrays;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;

public class IdentityAttributeRepositoryTest extends DBTest {

    @Autowired
    private IdentityAttributeRepository repository;

    @Autowired
    private EntityManager manager;

    @Test
    public void updateAssignableParameterTest() {
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

        var ids = Arrays.asList(val1.getId(), val2.getId());
        repository.updateAssignableParameter(ids, true);

        manager.flush();
        manager.clear();

        val1 = repository.getReferenceById(val1.getId());
        val2 = repository.getReferenceById(val2.getId());
        val3 = repository.getReferenceById(val3.getId());

        assertThat(val1.isAssignableToRoles()).isTrue();
        assertThat(val2.isAssignableToRoles()).isTrue();
        assertThat(val3.isAssignableToRoles()).isFalse();
    }
}
