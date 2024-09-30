package com.aruba.simpl.usersroles.repositories;

import static com.aruba.simpl.common.test.TestUtil.a;
import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.jupiter.params.provider.Arguments.arguments;

import com.aruba.simpl.usersroles.model.entities.IdentityAttributeRoles;
import com.aruba.simpl.usersroles.model.entities.Role;
import com.aruba.simpl.usersroles.utils.TransactionalUtils;
import jakarta.persistence.EntityManager;
import java.util.*;
import java.util.stream.Stream;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;
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

    @BeforeEach
    public void clearDB() {
        roleRepository.deleteAll();
        identityAttributeRolesRepository.deleteAll();
        manager.flush();
        manager.clear();
    }

    @Test
    public void loadContext() {}

    @Test
    public void updateRoleMappings() {

        var identityAttributeRoles = List.of(
                new IdentityAttributeRoles().setEnabled(true).setIdaCode("code_xyz"),
                new IdentityAttributeRoles().setEnabled(true).setIdaCode("code_abc"));

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
    public void findAlreadyAssignedToRole_withCommonValues_expectedValueInCommon() {

        var ias = List.of(iar("iar1"), iar("iar2"), iar("iar3"));

        identityAttributeRolesRepository.saveAll(ias);
        identityAttributeRolesRepository.save(iar("iar4"));

        var role1 = role("role1");
        var role2 = role("role2", ias);

        roleRepository.saveAll(List.of(role1, role2));

        var iasdb = identityAttributeRolesRepository.findAlreadyAssignedToRole("role2", List.of("iar3", "iar4"));
        assertThat(iasdb.size()).isEqualTo(1);
        assertThat(iasdb.get(0).getIdaCode()).isEqualTo("iar3");
    }

    @Test
    public void findAlreadyAssignedToRole_withoutCommonValues_expectedEmptyList() {

        var ias = List.of(iar("iar1"), iar("iar2"), iar("iar3"));

        identityAttributeRolesRepository.saveAll(ias);
        identityAttributeRolesRepository.save(iar("iar4"));

        var role1 = role("role1");
        var role2 = role("role2", ias);

        roleRepository.saveAll(List.of(role1, role2));

        var iasdb = identityAttributeRolesRepository.findAlreadyAssignedToRole("role2", List.of("iar4"));
        assertThat(iasdb.size()).isEqualTo(0);
    }

    public static Stream<Arguments> subtractIdentityAttributes_WithInputCombination_AsExpected() {
        return Stream.of(
                arguments(List.of("iar1", "iar2"), List.of("iar2"), List.of("iar1")),
                arguments(List.of("iar1", "iar2"), List.of("iar3"), List.of("iar1", "iar2")),
                arguments(List.of("iar1", "iar2"), List.of("iar1", "iar2"), List.of()),
                arguments(List.of("iar1", "iar2"), List.of("iar3", "iar4"), List.of("iar1", "iar2")),
                arguments(List.of("iar1", "iar2"), List.of(), List.of("iar1", "iar2")));
    }

    @ParameterizedTest
    @MethodSource
    public void subtractIdentityAttributes_WithInputCombination_AsExpected(
            List<String> iaCode1, List<String> iaCode2, List<String> iaCodeExpected) {

        var ias1 = iaCode1.stream().map(this::iar).toList();
        var ias2 = iaCode2.stream().map(this::iar).toList();

        identityAttributeRolesRepository.saveAll(ias1);
        identityAttributeRolesRepository.saveAll(ias2);

        var role1 = role("role1", ias1);
        var role2 = role("role2", ias2);

        roleRepository.saveAll(List.of(role1, role2));

        manager.flush();
        manager.clear();

        var iar = identityAttributeRolesRepository.subtractIdentityAttributes(role1, role2);

        assertThat(iar.stream().map(IdentityAttributeRoles::getIdaCode).toList())
                .containsAll(iaCodeExpected);
    }

    private Role role(String id) {
        return role(id, Collections.emptyList());
    }

    private Role role(String id, List<IdentityAttributeRoles> ia) {
        var role = a(Role.class);
        role.setId(id);
        role.setAssignedIdentityAttributes(ia);
        return role;
    }

    private IdentityAttributeRoles iar(String code) {
        var ia = a(IdentityAttributeRoles.class);
        ia.setId(null);
        ia.setIdaCode(code);
        return ia;
    }
}
