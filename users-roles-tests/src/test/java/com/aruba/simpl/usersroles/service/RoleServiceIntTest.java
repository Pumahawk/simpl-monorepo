package com.aruba.simpl.usersroles.service;

import static com.aruba.simpl.common.test.TestUtil.a;
import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.jupiter.params.provider.Arguments.arguments;

import com.aruba.simpl.usersroles.exceptions.IdentityAttributeAlreadyAssignedToRoleException;
import com.aruba.simpl.usersroles.exceptions.RoleNotFoundException;
import com.aruba.simpl.usersroles.model.entities.IdentityAttributeRoles;
import com.aruba.simpl.usersroles.model.entities.Role;
import com.aruba.simpl.usersroles.repositories.DBTest;
import com.aruba.simpl.usersroles.repositories.IdentityAttributeRolesRepository;
import com.aruba.simpl.usersroles.repositories.RoleRepository;
import jakarta.persistence.EntityManager;
import java.util.*;
import java.util.stream.Collectors;
import java.util.stream.Stream;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Import;

@Import({
    RoleService.class,
})
public class RoleServiceIntTest extends DBTest {
    @Autowired
    private RoleService roleService;

    @Autowired
    private IdentityAttributeRolesRepository identityAttributeRolesRepository;

    @Autowired
    private RoleRepository roleRepository;

    @Autowired
    private EntityManager manager;

    private Map<String, Role> roleDB;

    private Map<String, IdentityAttributeRoles> iarDB;

    @BeforeEach
    public void initDb() {
        identityAttributeRolesRepository.deleteAll();
        roleRepository.deleteAll();

        var ias = List.of(iar("ia1"), iar("ia2"), iar("ia3"));
        iarDB = ias.stream().collect(Collectors.toMap(IdentityAttributeRoles::getIdaCode, ia -> ia));

        identityAttributeRolesRepository.saveAll(ias);

        var roles = List.of(role("r1", iarDB.get("ia1"), iarDB.get("ia2")), role("r2", iarDB.get("ia3")), role("r3"));

        roleDB = roles.stream().collect(Collectors.toMap(Role::getId, r -> r));
        roleRepository.saveAll(roles);

        manager.flush();
        manager.clear();
    }

    public static Stream<Arguments> assignIdentityAttributes_matchAssignedAttribute() {
        return Stream.of(
                arguments("r1", List.of("ia3", "ia4"), List.of("ia1", "ia2", "ia3", "ia4")),
                arguments("r3", List.of("ia1", "ia4"), List.of("ia1", "ia4")),
                arguments("r2", List.of(), List.of("ia3")),
                arguments("r1", List.of("ia4"), List.of("ia1", "ia2", "ia4")));
    }

    @MethodSource
    @ParameterizedTest
    public void assignIdentityAttributes_matchAssignedAttribute(
            String roleId, List<String> assignedIAR, List<String> expectedIAR)
            throws IdentityAttributeAlreadyAssignedToRoleException, RoleNotFoundException {

        roleService.assignIdentityAttributes(roleId, assignedIAR);
        var r1 = roleRepository.findById(roleDB.get(roleId).getId()).orElseThrow();
        assertThat(r1.getAssignedIdentityAttributes().size()).isEqualTo(expectedIAR.size());
        var ias = r1.getAssignedIdentityAttributes().stream()
                .map(IdentityAttributeRoles::getIdaCode)
                .toList();
        assertThat(ias).contains(expectedIAR.toArray(new String[0]));
    }

    private Role role(String id, IdentityAttributeRoles... ia) {
        var role = a(Role.class);
        role.setId(id);
        role.setAssignedIdentityAttributes(Arrays.asList(ia));
        return role;
    }

    private IdentityAttributeRoles iar(String code) {
        var ia = a(IdentityAttributeRoles.class);
        ia.setId(null);
        ia.setIdaCode(code);
        return ia;
    }
}
