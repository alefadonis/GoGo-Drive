#!/bin/bash

BASE_URL="http://localhost:8081"
TEST_FILE="test_file.txt"
TEST_DIR="tests"
FILEPATH_TESTS="$TEST_DIR/$TEST_FILE"
DIV="--------------------------------------------------------------------------------"

check_status() {
  if [ $? -eq 0 ]; then
    echo -e "\n\033[1;32mPASS\033[0m"
  else
    echo -e "\n\033[1;31mFAIL\033[0m"
  fi
  echo $DIV
}

create_test_file() {
  echo "Criando arquivo de teste..."
  mkdir -p "$TEST_DIR"
  touch "$FILEPATH_TESTS"
  echo "Arquivo de teste criado em $TEST_DIR/$TEST_FILE."
  echo
}

test_upload_file() {
  echo "Teste UploadFile:"
  upload_url="$BASE_URL/upload"
  echo -e "URL do endpoint: \033[1m$upload_url\033[0m"
  curl -X POST -F "file=@$FILEPATH_TESTS" "$upload_url"
  check_status
}

test_download_file() {
  echo "Teste DownloadFile:"
  download_url="$BASE_URL/download/$TEST_FILE"
  echo -e "URL do endpoint: \033[1m$download_url\033[0m"
  cd tests/ && curl -OJ "$download_url"
  check_status
}

test_delete_file() {
  echo "Teste DeleteFile:"
  delete_url="$BASE_URL/delete/$TEST_FILE"
  echo -e "URL do endpoint: \033[1m$delete_url\033[0m"
  curl -X DELETE "$delete_url"
  check_status
}

test_list_files() {
  echo "Teste ListFiles:"
  list_files_url="$BASE_URL/files"
  echo -e "URL do endpoint: \033[1m$list_files_url\033[0m"
  curl "$list_files_url"
  echo
  check_status
}

echo -e "\033[1mIniciando execução dos testes:\033[0m"
echo $DIV

create_test_file
test_upload_file
test_download_file
test_delete_file
test_list_files