// Generated from /media/dados/workspace/2021/metric-query-parser/parser/Grammar.g4 by ANTLR 4.7.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class GrammarParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.7.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, TypeString=3, R_NEGATE=4, R_SIGLE_QUOTA=5, R_BRACE_L=6, 
		R_BRACE_R=7, R_BRACKET_L=8, R_BRACKET_R=9, R_PIPE=10, R_EQ=11, R_CONTAINS=12, 
		R_START_WITH=13, R_END_WITH=14, R_IN=15, R_GT=16, R_GTE=17, R_LT=18, R_LTE=19, 
		R_BETWEEN=20, R_NOT_BETWEEN=21, R_DOT=22, R_COMA=23, R_COLON=24, T_INTEIRO=25, 
		ID=26, WS=27;
	public static final int
		RULE_expression = 0, RULE_handler = 1, RULE_arguments = 2, RULE_argument = 3, 
		RULE_value = 4, RULE_typeHandler = 5, RULE_typeString = 6, RULE_typeBool = 7, 
		RULE_typeInt = 8, RULE_typeFloat = 9, RULE_operator = 10;
	public static final String[] ruleNames = {
		"expression", "handler", "arguments", "argument", "value", "typeHandler", 
		"typeString", "typeBool", "typeInt", "typeFloat", "operator"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "'true'", "'false'", null, "'!'", "'''", "'{'", "'}'", "'('", "')'", 
		"'|'", "'='", "'~='", "'|='", "'=|'", "'in'", "'>'", "'>='", "'<'", "'<='", 
		"'><'", "'<>'", "'.'", "','", "':'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, null, null, "TypeString", "R_NEGATE", "R_SIGLE_QUOTA", "R_BRACE_L", 
		"R_BRACE_R", "R_BRACKET_L", "R_BRACKET_R", "R_PIPE", "R_EQ", "R_CONTAINS", 
		"R_START_WITH", "R_END_WITH", "R_IN", "R_GT", "R_GTE", "R_LT", "R_LTE", 
		"R_BETWEEN", "R_NOT_BETWEEN", "R_DOT", "R_COMA", "R_COLON", "T_INTEIRO", 
		"ID", "WS"
	};
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Grammar.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public GrammarParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}
	public static class ExpressionContext extends ParserRuleContext {
		public HandlerContext handler() {
			return getRuleContext(HandlerContext.class,0);
		}
		public TerminalNode EOF() { return getToken(GrammarParser.EOF, 0); }
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		ExpressionContext _localctx = new ExpressionContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_expression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(22);
			handler();
			setState(23);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class HandlerContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode R_BRACKET_L() { return getToken(GrammarParser.R_BRACKET_L, 0); }
		public TerminalNode R_BRACKET_R() { return getToken(GrammarParser.R_BRACKET_R, 0); }
		public ArgumentsContext arguments() {
			return getRuleContext(ArgumentsContext.class,0);
		}
		public HandlerContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_handler; }
	}

	public final HandlerContext handler() throws RecognitionException {
		HandlerContext _localctx = new HandlerContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_handler);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(25);
			match(ID);
			setState(26);
			match(R_BRACKET_L);
			setState(28);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ID) {
				{
				setState(27);
				arguments();
				}
			}

			setState(30);
			match(R_BRACKET_R);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ArgumentsContext extends ParserRuleContext {
		public List<ArgumentContext> argument() {
			return getRuleContexts(ArgumentContext.class);
		}
		public ArgumentContext argument(int i) {
			return getRuleContext(ArgumentContext.class,i);
		}
		public List<TerminalNode> R_COMA() { return getTokens(GrammarParser.R_COMA); }
		public TerminalNode R_COMA(int i) {
			return getToken(GrammarParser.R_COMA, i);
		}
		public ArgumentsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arguments; }
	}

	public final ArgumentsContext arguments() throws RecognitionException {
		ArgumentsContext _localctx = new ArgumentsContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_arguments);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(32);
			argument();
			setState(37);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==R_COMA) {
				{
				{
				setState(33);
				match(R_COMA);
				setState(34);
				argument();
				}
				}
				setState(39);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ArgumentContext extends ParserRuleContext {
		public HandlerContext handler() {
			return getRuleContext(HandlerContext.class,0);
		}
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public OperatorContext operator() {
			return getRuleContext(OperatorContext.class,0);
		}
		public ValueContext value() {
			return getRuleContext(ValueContext.class,0);
		}
		public ArgumentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_argument; }
	}

	public final ArgumentContext argument() throws RecognitionException {
		ArgumentContext _localctx = new ArgumentContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_argument);
		try {
			setState(45);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(40);
				handler();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(41);
				match(ID);
				setState(42);
				operator();
				setState(43);
				value();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ValueContext extends ParserRuleContext {
		public TypeStringContext typeString() {
			return getRuleContext(TypeStringContext.class,0);
		}
		public TypeBoolContext typeBool() {
			return getRuleContext(TypeBoolContext.class,0);
		}
		public TypeIntContext typeInt() {
			return getRuleContext(TypeIntContext.class,0);
		}
		public TypeFloatContext typeFloat() {
			return getRuleContext(TypeFloatContext.class,0);
		}
		public TypeHandlerContext typeHandler() {
			return getRuleContext(TypeHandlerContext.class,0);
		}
		public ValueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_value; }
	}

	public final ValueContext value() throws RecognitionException {
		ValueContext _localctx = new ValueContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_value);
		try {
			setState(52);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(47);
				typeString();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(48);
				typeBool();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(49);
				typeInt();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(50);
				typeFloat();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(51);
				typeHandler();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeHandlerContext extends ParserRuleContext {
		public HandlerContext handler() {
			return getRuleContext(HandlerContext.class,0);
		}
		public TypeHandlerContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeHandler; }
	}

	public final TypeHandlerContext typeHandler() throws RecognitionException {
		TypeHandlerContext _localctx = new TypeHandlerContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_typeHandler);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(54);
			handler();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeStringContext extends ParserRuleContext {
		public TerminalNode TypeString() { return getToken(GrammarParser.TypeString, 0); }
		public TypeStringContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeString; }
	}

	public final TypeStringContext typeString() throws RecognitionException {
		TypeStringContext _localctx = new TypeStringContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_typeString);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(56);
			match(TypeString);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeBoolContext extends ParserRuleContext {
		public TypeBoolContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeBool; }
	}

	public final TypeBoolContext typeBool() throws RecognitionException {
		TypeBoolContext _localctx = new TypeBoolContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_typeBool);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(58);
			_la = _input.LA(1);
			if ( !(_la==T__0 || _la==T__1) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeIntContext extends ParserRuleContext {
		public TerminalNode T_INTEIRO() { return getToken(GrammarParser.T_INTEIRO, 0); }
		public TypeIntContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeInt; }
	}

	public final TypeIntContext typeInt() throws RecognitionException {
		TypeIntContext _localctx = new TypeIntContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_typeInt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(60);
			match(T_INTEIRO);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeFloatContext extends ParserRuleContext {
		public List<TerminalNode> T_INTEIRO() { return getTokens(GrammarParser.T_INTEIRO); }
		public TerminalNode T_INTEIRO(int i) {
			return getToken(GrammarParser.T_INTEIRO, i);
		}
		public TerminalNode R_DOT() { return getToken(GrammarParser.R_DOT, 0); }
		public TypeFloatContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeFloat; }
	}

	public final TypeFloatContext typeFloat() throws RecognitionException {
		TypeFloatContext _localctx = new TypeFloatContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_typeFloat);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(62);
			match(T_INTEIRO);
			setState(63);
			match(R_DOT);
			setState(64);
			match(T_INTEIRO);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class OperatorContext extends ParserRuleContext {
		public TerminalNode R_NEGATE() { return getToken(GrammarParser.R_NEGATE, 0); }
		public OperatorContext operator() {
			return getRuleContext(OperatorContext.class,0);
		}
		public TerminalNode R_EQ() { return getToken(GrammarParser.R_EQ, 0); }
		public TerminalNode R_CONTAINS() { return getToken(GrammarParser.R_CONTAINS, 0); }
		public TerminalNode R_START_WITH() { return getToken(GrammarParser.R_START_WITH, 0); }
		public TerminalNode R_END_WITH() { return getToken(GrammarParser.R_END_WITH, 0); }
		public TerminalNode R_IN() { return getToken(GrammarParser.R_IN, 0); }
		public TerminalNode R_GT() { return getToken(GrammarParser.R_GT, 0); }
		public TerminalNode R_GTE() { return getToken(GrammarParser.R_GTE, 0); }
		public TerminalNode R_LT() { return getToken(GrammarParser.R_LT, 0); }
		public TerminalNode R_LTE() { return getToken(GrammarParser.R_LTE, 0); }
		public TerminalNode R_BETWEEN() { return getToken(GrammarParser.R_BETWEEN, 0); }
		public TerminalNode R_NOT_BETWEEN() { return getToken(GrammarParser.R_NOT_BETWEEN, 0); }
		public OperatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operator; }
	}

	public final OperatorContext operator() throws RecognitionException {
		OperatorContext _localctx = new OperatorContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_operator);
		try {
			setState(79);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case R_NEGATE:
				enterOuterAlt(_localctx, 1);
				{
				setState(66);
				match(R_NEGATE);
				setState(67);
				operator();
				}
				break;
			case R_EQ:
				enterOuterAlt(_localctx, 2);
				{
				setState(68);
				match(R_EQ);
				}
				break;
			case R_CONTAINS:
				enterOuterAlt(_localctx, 3);
				{
				setState(69);
				match(R_CONTAINS);
				}
				break;
			case R_START_WITH:
				enterOuterAlt(_localctx, 4);
				{
				setState(70);
				match(R_START_WITH);
				}
				break;
			case R_END_WITH:
				enterOuterAlt(_localctx, 5);
				{
				setState(71);
				match(R_END_WITH);
				}
				break;
			case R_IN:
				enterOuterAlt(_localctx, 6);
				{
				setState(72);
				match(R_IN);
				}
				break;
			case R_GT:
				enterOuterAlt(_localctx, 7);
				{
				setState(73);
				match(R_GT);
				}
				break;
			case R_GTE:
				enterOuterAlt(_localctx, 8);
				{
				setState(74);
				match(R_GTE);
				}
				break;
			case R_LT:
				enterOuterAlt(_localctx, 9);
				{
				setState(75);
				match(R_LT);
				}
				break;
			case R_LTE:
				enterOuterAlt(_localctx, 10);
				{
				setState(76);
				match(R_LTE);
				}
				break;
			case R_BETWEEN:
				enterOuterAlt(_localctx, 11);
				{
				setState(77);
				match(R_BETWEEN);
				}
				break;
			case R_NOT_BETWEEN:
				enterOuterAlt(_localctx, 12);
				{
				setState(78);
				match(R_NOT_BETWEEN);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\35T\4\2\t\2\4\3\t"+
		"\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t\13\4"+
		"\f\t\f\3\2\3\2\3\2\3\3\3\3\3\3\5\3\37\n\3\3\3\3\3\3\4\3\4\3\4\7\4&\n\4"+
		"\f\4\16\4)\13\4\3\5\3\5\3\5\3\5\3\5\5\5\60\n\5\3\6\3\6\3\6\3\6\3\6\5\6"+
		"\67\n\6\3\7\3\7\3\b\3\b\3\t\3\t\3\n\3\n\3\13\3\13\3\13\3\13\3\f\3\f\3"+
		"\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\5\fR\n\f\3\f\2\2\r\2\4\6\b"+
		"\n\f\16\20\22\24\26\2\3\3\2\3\4\2Z\2\30\3\2\2\2\4\33\3\2\2\2\6\"\3\2\2"+
		"\2\b/\3\2\2\2\n\66\3\2\2\2\f8\3\2\2\2\16:\3\2\2\2\20<\3\2\2\2\22>\3\2"+
		"\2\2\24@\3\2\2\2\26Q\3\2\2\2\30\31\5\4\3\2\31\32\7\2\2\3\32\3\3\2\2\2"+
		"\33\34\7\34\2\2\34\36\7\n\2\2\35\37\5\6\4\2\36\35\3\2\2\2\36\37\3\2\2"+
		"\2\37 \3\2\2\2 !\7\13\2\2!\5\3\2\2\2\"\'\5\b\5\2#$\7\31\2\2$&\5\b\5\2"+
		"%#\3\2\2\2&)\3\2\2\2\'%\3\2\2\2\'(\3\2\2\2(\7\3\2\2\2)\'\3\2\2\2*\60\5"+
		"\4\3\2+,\7\34\2\2,-\5\26\f\2-.\5\n\6\2.\60\3\2\2\2/*\3\2\2\2/+\3\2\2\2"+
		"\60\t\3\2\2\2\61\67\5\16\b\2\62\67\5\20\t\2\63\67\5\22\n\2\64\67\5\24"+
		"\13\2\65\67\5\f\7\2\66\61\3\2\2\2\66\62\3\2\2\2\66\63\3\2\2\2\66\64\3"+
		"\2\2\2\66\65\3\2\2\2\67\13\3\2\2\289\5\4\3\29\r\3\2\2\2:;\7\5\2\2;\17"+
		"\3\2\2\2<=\t\2\2\2=\21\3\2\2\2>?\7\33\2\2?\23\3\2\2\2@A\7\33\2\2AB\7\30"+
		"\2\2BC\7\33\2\2C\25\3\2\2\2DE\7\6\2\2ER\5\26\f\2FR\7\r\2\2GR\7\16\2\2"+
		"HR\7\17\2\2IR\7\20\2\2JR\7\21\2\2KR\7\22\2\2LR\7\23\2\2MR\7\24\2\2NR\7"+
		"\25\2\2OR\7\26\2\2PR\7\27\2\2QD\3\2\2\2QF\3\2\2\2QG\3\2\2\2QH\3\2\2\2"+
		"QI\3\2\2\2QJ\3\2\2\2QK\3\2\2\2QL\3\2\2\2QM\3\2\2\2QN\3\2\2\2QO\3\2\2\2"+
		"QP\3\2\2\2R\27\3\2\2\2\7\36\'/\66Q";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}