package contact

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/contact"
)

func trace(span trace1.Span, in *npool.ContactReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("UsedFor.%v", index), in.GetUsedFor().String()),
		attribute.String(fmt.Sprintf("Account.%v", index), in.GetAccount()),
		attribute.String(fmt.Sprintf("AccountType.%v", index), in.GetAccountType().String()),
		attribute.String(fmt.Sprintf("Sender.%v", index), in.GetSender()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.ContactReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Val", in.GetID().GetValue()),
		attribute.String("AppID.Op", in.GetAppID().GetOp()),
		attribute.String("AppID.Val", in.GetAppID().GetValue()),
		attribute.String("UsedFor.Op", in.GetUsedFor().GetOp()),
		attribute.Int("UsedFor.Val", int(in.GetUsedFor().GetValue())),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.ContactReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
