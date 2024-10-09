package com.aruba.simpl.usersroles.service;

import com.aruba.simpl.common.exchanges.mtls.AuthorityExchange;
import com.aruba.simpl.common.security.JwtService;
import com.aruba.simpl.usersroles.configurations.MtlsClientBuilder;
import com.aruba.simpl.usersroles.services.CredentialService;
import com.aruba.simpl.usersroles.services.EphemeralProofService;
import com.aruba.simpl.usersroles.services.IdentityAttributeService;
import com.aruba.simpl.usersroles.services.impl.AgentServiceImpl;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.context.annotation.Import;
import org.springframework.test.context.junit.jupiter.SpringExtension;

@ExtendWith(SpringExtension.class)
@Import({
    AgentServiceImpl.class,
})
@MockBean({
    AuthorityExchange.class,
    IdentityAttributeService.class,
    EphemeralProofService.class,
    JwtService.class,
    CredentialService.class,
    MtlsClientBuilder.class,
})
public class AgentServiceIntTest {

    @Test
    public void loadContext() {}
}
