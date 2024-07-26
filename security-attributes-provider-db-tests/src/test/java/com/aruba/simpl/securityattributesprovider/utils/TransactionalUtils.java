package com.aruba.simpl.securityattributesprovider.utils;

import org.springframework.stereotype.Component;
import org.springframework.transaction.annotation.Transactional;

@Component
public class TransactionalUtils {

    @Transactional
    public void transactional(Runnable fn) {
        fn.run();
    }
}
