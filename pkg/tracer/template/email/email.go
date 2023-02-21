package email

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/template/email"
)

func trace(span trace1.Span, in *npool.EmailTemplateReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("LangID.%v", index), in.GetLangID()),
		attribute.String(fmt.Sprintf("DefaultToUsername.%v", index), in.GetDefaultToUsername()),
		attribute.String(fmt.Sprintf("UsedFor.%v", index), in.GetUsedFor().String()),
		attribute.String(fmt.Sprintf("Sender.%v", index), in.GetSender()),
		attribute.StringSlice(fmt.Sprintf("ReplyTos.%v", index), in.GetReplyTos()),
		attribute.StringSlice(fmt.Sprintf("CCTos.%v", index), in.GetCCTos()),
		attribute.String(fmt.Sprintf("Subject.%v", index), in.GetSubject()),
		attribute.String(fmt.Sprintf("Body.%v", index), in.GetBody()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.EmailTemplateReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Val", in.GetID().GetValue()),
		attribute.String("AppID.Op", in.GetAppID().GetOp()),
		attribute.String("AppID.Val", in.GetAppID().GetValue()),
		attribute.String("LangID.Op", in.GetLangID().GetOp()),
		attribute.String("LangID.Val", in.GetLangID().GetValue()),
		attribute.String("Sender.Op", in.GetSender().GetOp()),
		attribute.String("Sender.Val", in.GetSender().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.EmailTemplateReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
