package com.example.hls.controller;

import com.example.hls.service.HlsService;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.core.io.Resource;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

import java.io.IOException;

@RestController
@RequiredArgsConstructor
@Slf4j
@CrossOrigin(origins = "*")
public class HlsController {

    private final HlsService hlsService;

    @GetMapping("/{hlsName}.m3u8")
    public ResponseEntity<Resource> getHls(@PathVariable String hlsName) throws IOException {
        Resource resource = hlsService.getVideoRes(hlsName + ".m3u8");

        HttpHeaders headers = new HttpHeaders();
        headers.set(HttpHeaders.CONTENT_DISPOSITION, "attachment; filename=playlist.m3u8");
        headers.setContentType(MediaType.parseMediaType("application/vnd.apple.mpegurl"));
        return new ResponseEntity<>(resource, headers, HttpStatus.OK);
    }

    @GetMapping("/{tsName}.ts")
    public ResponseEntity<Resource> getHlsTs(@PathVariable String tsName) throws IOException {

        tsName = tsName + ".ts";
        Resource resource = hlsService.getVideoRes(tsName);

        HttpHeaders headers = new HttpHeaders();
        headers.set(HttpHeaders.CONTENT_DISPOSITION, "attachment; filename=" + tsName);
        headers.setContentType(MediaType.parseMediaType(MediaType.APPLICATION_OCTET_STREAM_VALUE));
        return new ResponseEntity<>(resource, headers, HttpStatus.OK);
    }

    @GetMapping("/file/{fileName}")
    public ResponseEntity<Resource> getHlsFile(@PathVariable String fileName) throws IOException {
        Resource resource = hlsService.getVideoRes(fileName);
        return new ResponseEntity<>(resource, HttpStatus.OK);
    }

    @GetMapping("/convert")
    public ResponseEntity<?> convert() throws IOException {
        hlsService.convertHls();
        return ResponseEntity.ok().build();
    }
}
