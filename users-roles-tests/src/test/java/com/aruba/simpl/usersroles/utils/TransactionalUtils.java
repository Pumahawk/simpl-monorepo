package com.aruba.simpl.usersroles.utils;

import org.springframework.stereotype.Component;
import org.springframework.transaction.annotation.Transactional;

@Component
public class TransactionalUtils {

    @Transactional
    public void transactional(Runnable fn) {
        fn.run();
    }
}
