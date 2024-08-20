package com.aruba.simpl.usersroles.repositories;

import com.aruba.simpl.usersroles.model.entities.IdentityAttributeRoles;
import com.aruba.simpl.usersroles.model.entities.Role;
import java.util.List;
import java.util.UUID;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;

public class RoleRepositoryTest extends DBTest {

    @Autowired
    private IdentityAttributeRolesRepository identityAttributeRolesRepository;

    @Autowired
    private RoleRepository roleRepository;

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
}
