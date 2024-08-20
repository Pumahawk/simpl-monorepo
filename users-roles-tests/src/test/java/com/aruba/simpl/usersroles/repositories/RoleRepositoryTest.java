package com.aruba.simpl.usersroles.repositories;

import static com.aruba.simpl.common.test.TestUtil.a;
import static org.assertj.core.api.Assertions.assertThat;

import com.aruba.simpl.usersroles.model.entities.IdentityAttributeRoles;
import com.aruba.simpl.usersroles.model.entities.Role;
import com.aruba.simpl.usersroles.utils.TransactionalUtils;
import jakarta.persistence.EntityManager;
import java.util.ArrayList;
import java.util.List;
import java.util.UUID;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;

public class RoleRepositoryTest extends DBTest {

    @Autowired
    private IdentityAttributeRolesRepository identityAttributeRolesRepository;

    @Autowired
    private RoleRepository roleRepository;

    @Autowired
    TransactionalUtils transactional;

    @Autowired
    private EntityManager manager;

    @Test
    public void loadContext() {}

    @Test
    public void updateRoleMappings() {

        var identityAttributeRoles = List.of(
                new IdentityAttributeRoles()
                        .setId(UUID.randomUUID())
                        .setEnabled(true)
                        .setIdaCode("code_xyz"),
                new IdentityAttributeRoles()
                        .setId(UUID.randomUUID())
                        .setEnabled(true)
                        .setIdaCode("code_abc"));

        var role = new Role()
                .setId(UUID.randomUUID().toString())
                .setName("unit-tests-role")
                .setDescr("Unit tests description")
                .setAssignedIdentityAttributes(identityAttributeRoles);
        role.setAssignedIdentityAttributes(identityAttributeRoles);
        identityAttributeRolesRepository.saveAll(identityAttributeRoles);
        roleRepository.save(role);
    }

    @Test
    public void findAlreadyAssignedToRole_extractFrom_expectedSuccess() {

        var ias = List.of(iar("iar1"), iar("iar2"), iar("iar3"));

        identityAttributeRolesRepository.saveAll(ias);
        identityAttributeRolesRepository.save(iar("iar4"));

        var role1 = role("role1");
        var role2 = role("role2");
        role2.getAssignedIdentityAttributes().addAll(ias);

        roleRepository.saveAll(List.of(role1, role2));

        var iasdb = identityAttributeRolesRepository.findAlreadyAssignedToRole("role2", List.of("iar3", "iar4"));
        assertThat(iasdb.size()).isEqualTo(1);
        assertThat(iasdb.get(0).getIdaCode()).isEqualTo("iar3");
    }

    private static Role role(String id) {
        var role = a(Role.class);
        role.setId(id);
        role.setAssignedIdentityAttributes(new ArrayList<>());
        return role;
    }

    private IdentityAttributeRoles iar(String idaCode) {
        var ia = a(IdentityAttributeRoles.class);
        ia.setEnabled(true);
        ia.setIdaCode(idaCode);
        ia.setId(UUID.randomUUID());
        return ia;
    }
}
