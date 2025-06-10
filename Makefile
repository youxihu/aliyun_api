.PHONY: build

# 默认 build 提示信息
build:
	@echo "需要指定构建编号"
	@echo "1: 查询安全组内规则信息"
	@echo "1: 查询安全组内规则描述带*的信息"
	@echo "3: 修改安全组入方向为允许公司网络访问"
	@echo "例如: make build-1"

# 编号构建逻辑
build-%:
	@case '$*' in \
		1) APP_NAME=sgAttr ;; \
		2) APP_NAME=sgStarred ;; \
		3) APP_NAME=sgModify ;; \
		*) echo "无效构建编号" ; exit 1 ;; \
	esac && \
	mkdir -p bin/ && \
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/$$APP_NAME cmd/$$APP_NAME/main.go && \
	upx ./bin/$$APP_NAME


test:
	@echo "需要指定测试编号"
	@echo "1: 测试查询安全组内规则信息 "
	@echo "shell: go test -v -run TestDescribeSecurityGroupAttribute internal/describeSecurityGroups"
	@echo "2: 测试查询安全组内规则描述带*的信息"
	@echo "shell: go test -v -run TestGetStarredRules internal/describeSecurityGroups"
	@echo "3: 测试修改安全组入方向为允许公司网络访问"
	@echo "shell: go test -v -run TestModifySecurityGroup internal/modifySecurityGroups"
	@echo "例如:  make test-1"


test-%:
	@case '$*' in \
		1) cd internal/describeSecurityGroups && go test -v -run TestDescribeSecurityGroupAttribute ;; \
		2) cd internal/describeSecurityGroups  && go test -v -run TestGetStarredRules ;; \
		3) cd internal/modifySecurityGroups  && go test -v -run TestModifySecurityGroup ;; \
		*) echo "无效选项" ;; \
	esac