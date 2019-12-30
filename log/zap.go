package log

import (
	"fmt"
	"github.com/uasouz/tms/config"
	"go.uber.org/zap"
	"strings"
)

var logger *zap.SugaredLogger

// Operation a operacao usada na API
var Operation string

// Recipient o nome do banco
var Recipient string

// Log struct com os elemtos do log
type Log struct {
	Operation   string
	Recipient   string
	RequestKey  string
	BankName    string
	IPAddress   string
	NossoNumero uint
	Logger      *zap.SugaredLogger
}

func CreateLog() *Log {
	_logger, _ := zap.NewProduction()
	logger = _logger.Sugar()
	return &Log{
		Logger: logger,
	}
}

func formatter(message string) string {
	return "[{Application}: {Operation}] - {MessageType} " + message
}

// Request loga o request para algum banco
func (l Log) Request(content interface{}, url string, headers map[string]string) {
	if config.Get().DisableLog {
		return
	}
	go (func() {
		//props := l.defaultProperties("Request", content)
		//props.AddProperty("Headers", headers)
		//props.AddProperty("URL", url)
		action := strings.Split(url, "/")
		msg := formatter(fmt.Sprintf("to {BankName} (%s) | {Recipient}", action[len(action)-1]))
		//
		//l.logger.Information(msg, props)
		l.Logger.Infow(msg,
			"Request", content,
			"URL", url,
			"Headers", headers,
		)
	})()
}

// Response loga o response para algum banco
func (l Log) Response(content interface{}, url string) {
	if config.Get().DisableLog {
		return
	}
	go (func() {
		//props := l.defaultProperties("Response", content)
		//props.AddProperty("URL", url)
		action := strings.Split(url, "/")
		msg := formatter(fmt.Sprintf("from {BankName} (%s) | {Recipient}", action[len(action)-1]))
		l.Logger.Infow(msg,
			"Response", content,
			"URL", url,
		)
		//l.logger.Information(msg, props)
	})()
}

//Info loga mensagem do level INFO
func (l Log) Info(msg string) {
	if config.Get().DisableLog {
		return
	}
	go logger.Infow(msg)
}

func Info(msg string) {
	if config.Get().DisableLog {
		return
	}
	go logger.Infow(msg)
}

//Warn loga mensagem do leve Warning
func (l Log) Warn(content interface{}, msg string) {
	if config.Get().DisableLog {
		return
	}
	go (func() {
		m := formatter(msg)

		l.Logger.Warnw(m, "Warning", content)
	})()
}

// Fatal loga erros da aplicação
func (l Log) Fatal(content interface{}, msg string) {
	if config.Get().DisableLog {
		return
	}
	go (func() {
		m := formatter(msg)

		l.Logger.Fatal(m, "Error", content)
	})()
}
