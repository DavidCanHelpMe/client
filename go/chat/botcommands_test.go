package chat

import (
	"testing"
	"time"

	"github.com/keybase/client/go/chat/types"
	"github.com/keybase/client/go/chat/utils"
	"github.com/keybase/client/go/protocol/chat1"
	"github.com/keybase/client/go/protocol/gregor1"
	"github.com/stretchr/testify/require"
)

func TestBotCommandManager(t *testing.T) {
	useRemoteMock = false
	defer func() { useRemoteMock = true }()
	ctc := makeChatTestContext(t, "TestBotCommandManager", 2)
	defer ctc.cleanup()

	timeout := 20 * time.Second
	users := ctc.users()
	tc := ctc.world.Tcs[users[0].Username]
	tc1 := ctc.world.Tcs[users[1].Username]
	ctx := ctc.as(t, users[0]).startCtx
	ctx1 := ctc.as(t, users[1]).startCtx
	uid := gregor1.UID(users[0].GetUID().ToBytes())
	t.Logf("uid: %s", uid)
	listener0 := newServerChatListener()
	ctc.as(t, users[0]).h.G().NotifyRouter.AddListener(listener0)

	impConv := mustCreateConversationForTest(t, ctc, users[0], chat1.TopicType_CHAT,
		chat1.ConversationMembersType_IMPTEAMNATIVE)
	impConv1 := mustCreateConversationForTest(t, ctc, users[1], chat1.TopicType_CHAT,
		chat1.ConversationMembersType_IMPTEAMNATIVE, users[0])
	t.Logf("impconv: %s", impConv.Id)
	t.Logf("impconv1: %s", impConv1.Id)
	teamConv := mustCreateConversationForTest(t, ctc, users[0], chat1.TopicType_CHAT,
		chat1.ConversationMembersType_TEAM)
	t.Logf("teamconv: %x", teamConv.Id.DbShortForm())

	// test public
	alias := "MIKE BOT"
	commands := []chat1.AdvertiseCommandsParam{
		chat1.AdvertiseCommandsParam{
			Typ: chat1.BotCommandsAdvertisementTyp_PUBLIC,
			Commands: []chat1.UserBotCommandInput{
				chat1.UserBotCommandInput{
					Name:        "status",
					Description: "get status",
					Usage:       "just type it",
				},
			},
		},
	}
	require.NoError(t, tc.Context().BotCommandManager.Advertise(ctx, &alias, commands))
	cmds, err := tc.Context().BotCommandManager.ListCommands(ctx, impConv.Id)
	require.NoError(t, err)
	require.Zero(t, len(cmds))
	errCh, err := tc.Context().BotCommandManager.UpdateCommands(ctx, impConv.Id, nil)
	require.NoError(t, err)
	errCh1, err := tc1.Context().BotCommandManager.UpdateCommands(ctx1, impConv1.Id, nil)
	require.NoError(t, err)

	select {
	case updates := <-listener0.threadsStale:
		require.Equal(t, 1, len(updates))
		require.Equal(t, impConv.Id, updates[0].ConvID)
		require.Equal(t, chat1.StaleUpdateType_CONVUPDATE, updates[0].UpdateType)
	case <-time.After(timeout):
		require.Fail(t, "no stale")
	}
	impConvLocal, err := utils.GetVerifiedConv(ctx, tc.Context(), uid, impConv.Id,
		types.InboxSourceDataSourceAll)
	require.NoError(t, err)
	typ, err := impConvLocal.BotCommands.Typ()
	require.NoError(t, err)
	require.Equal(t, chat1.ConversationCommandGroupsTyp_CUSTOM, typ)
	require.Equal(t, 1, len(impConvLocal.BotCommands.Custom().Commands))
	require.Equal(t, "status", impConvLocal.BotCommands.Custom().Commands[0].Name)
	require.NoError(t, <-errCh)
	cmds, err = tc.Context().BotCommandManager.ListCommands(ctx, impConv.Id)
	require.NoError(t, err)
	require.Equal(t, 1, len(cmds))
	require.Equal(t, "status", cmds[0].Name)
	require.NoError(t, <-errCh1)
	cmds, err = tc1.Context().BotCommandManager.ListCommands(ctx1, impConv1.Id)
	require.NoError(t, err)
	require.Equal(t, 1, len(cmds))
	require.Equal(t, "status", cmds[0].Name)

	// test team
	commands = append(commands, chat1.AdvertiseCommandsParam{
		Typ: chat1.BotCommandsAdvertisementTyp_TLFID_CONVS,
		Commands: []chat1.UserBotCommandInput{chat1.UserBotCommandInput{
			Name: "teamconvonly",
		}},
		TeamName: &teamConv.TlfName,
	}, chat1.AdvertiseCommandsParam{
		Typ: chat1.BotCommandsAdvertisementTyp_TLFID_MEMBERS,
		Commands: []chat1.UserBotCommandInput{chat1.UserBotCommandInput{
			Name: "teammembsonly",
		}},
		TeamName: &teamConv.TlfName,
	})
	require.NoError(t, tc.Context().BotCommandManager.Advertise(ctx, &alias, commands))
	errCh, err = tc.Context().BotCommandManager.UpdateCommands(ctx, impConv.Id, nil)
	require.NoError(t, err)
	errChT, err := tc.Context().BotCommandManager.UpdateCommands(ctx, teamConv.Id, nil)
	require.NoError(t, err)
	errCh1, err = tc.Context().BotCommandManager.UpdateCommands(ctx, impConv.Id, nil)
	require.NoError(t, err)
	require.NoError(t, <-errCh)
	require.NoError(t, <-errCh1)
	require.NoError(t, <-errChT)
	cmds, err = tc.Context().BotCommandManager.ListCommands(ctx, impConv.Id)
	require.NoError(t, err)
	require.Equal(t, 2, len(cmds))
	cmds, err = tc.Context().BotCommandManager.ListCommands(ctx, teamConv.Id)
	require.NoError(t, err)
	require.Equal(t, 3, len(cmds))
	cmds, err = tc1.Context().BotCommandManager.ListCommands(ctx1, impConv1.Id)
	require.NoError(t, err)
	require.Equal(t, 1, len(cmds))
}