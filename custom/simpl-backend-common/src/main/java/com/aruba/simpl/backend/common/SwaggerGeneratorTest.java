package com.aruba.simpl.backend.common;

import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

import java.io.ByteArrayInputStream;
import java.io.FileOutputStream;
import org.apache.commons.compress.utils.IOUtils;
import org.junit.jupiter.api.Test;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.test.web.servlet.MockMvc;

public class SwaggerGeneratorTest {

    Logger logger = LoggerFactory.getLogger(getClass());

    @Autowired
    MockMvc mockMvc;

    @Test
    public void generateJson() throws Exception {
        mockMvc.perform(get("/v3/api-docs")).andExpect(status().is(200)).andDo(result -> {
            var outputPath = getClass().getResource("/").getPath() + "swagger.json";
            logger.info("Output swagger path: {}", outputPath);
            IOUtils.copy(
                    new ByteArrayInputStream(result.getResponse().getContentAsByteArray()),
                    new FileOutputStream(outputPath));
        });
    }
}
