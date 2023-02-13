package notif

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif"
)

func trace(span trace1.Span, in *npool.NotifReq, index int) trace1.Span {
	var channel []string
	for _, val := range in.GetChannels() {
		channel = append(channel, val.String())
	}
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("UserID.%v", index), in.GetUserID()),
		attribute.Bool(fmt.Sprintf("AlreadyRead.%v", index), in.GetAlreadyRead()),
		attribute.String(fmt.Sprintf("LangID.%v", index), in.GetLangID()),
		attribute.String(fmt.Sprintf("EventType.%v", index), in.GetEventType().String()),
		attribute.Bool(fmt.Sprintf("UseTemplate.%v", index), in.GetUseTemplate()),
		attribute.String(fmt.Sprintf("Title.%v", index), in.GetTitle()),
		attribute.String(fmt.Sprintf("Content.%v", index), in.GetContent()),
		attribute.StringSlice(fmt.Sprintf("Channels.%v", index), channel),
		attribute.Bool(fmt.Sprintf("EmailSend.%v", index), in.GetEmailSend()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.NotifReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Value", in.GetID().GetValue()),
		attribute.String("AppID.Op", in.GetAppID().GetOp()),
		attribute.String("AppID.Value", in.GetAppID().GetValue()),
		attribute.String("UserID.Op", in.GetUserID().GetOp()),
		attribute.String("UserID.Value", in.GetUserID().GetValue()),
		attribute.String("AlreadyRead.Op", in.GetAlreadyRead().GetOp()),
		attribute.Bool("AlreadyRead.Value", in.GetAlreadyRead().GetValue()),
		attribute.String("LangID.Op", in.GetLangID().GetOp()),
		attribute.String("LangID.Value", in.GetLangID().GetValue()),
		attribute.String("EventType.Op", in.GetEventType().GetOp()),
		attribute.Int("EventType.Value", int(in.GetEventType().GetValue())),
		attribute.String("UseTemplate.Op", in.GetUseTemplate().GetOp()),
		attribute.Bool("UseTemplate.Value", in.GetUseTemplate().GetValue()),
		attribute.String("Channels.Op", in.GetChannels().GetOp()),
		attribute.StringSlice("Channels.Value", in.GetChannels().GetValue()),
		attribute.String("EmailSend.Op", in.GetEmailSend().GetOp()),
		attribute.Bool("EmailSend.Value", in.GetEmailSend().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.NotifReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
