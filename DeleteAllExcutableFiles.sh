#!/bin/bash

# Удаляем все исполняемые файлы, кроме .sh скриптов
echo "Удаление Go-бинарников..."

# Находим и удаляем статически линкованные ELF файлы (это Go-бинарники)
find . -type f -executable ! -name "*.sh" -exec sh -c '
    for file do
        if file "$file" | grep -q "statically linked"; then
            echo "Удаляю: $file"
            rm -f "$file"
        fi
    done
' _ {} +

echo "Готово"
