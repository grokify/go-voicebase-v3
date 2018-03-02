java -jar swagger-codegen-cli.jar generate -c swagger_codegen_config.json -i swagger_spec_voicebase_v3.yaml -l go -o ./client
perl -p -i -e 's/(\S+\s+VbIncludeTypeEnum\s+=\s+")/INCLUDE_$1/g' client/vb_include_type_enum.go
perl -p -i -e 's/(\S+\s+VbTaskStatusEnum\s+=\s+")/TASK_STATUS_$1/g' client/vb_task_status_enum.go
perl -p -i -e 's/\*(VbVocabularyType)/$1/g' client/vb_vocabulary.go
gofmt -s -w client/*.go