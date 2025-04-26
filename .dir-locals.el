((go-ts-mode
  . ((eval
      . (let ((root (expand-file-name (project-root (project-current)))))
          (setq-local eglot-server-programs
                      `((go-ts-mode
                         "podman" "run" "--rm" "--interactive"
                         ,(concat "--volume=" root ":" root ":z")
                         ,(concat "--workdir=" root)
                         "ghcr.io/bkhl/image-gopls:latest")))
          (add-hook 'before-save-hook #'eglot-format-buffer nil t)
          (eglot-ensure))))))
