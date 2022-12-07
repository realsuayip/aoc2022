package day7

import (
	"fmt"
	"strconv"
	"strings"
)

var input = "$ cd /\n$ ls\ndir jmtrrrp\ndir jssnn\ndir lbrmb\n11968 pcccp\n$ cd jmtrrrp\n$ ls\n77968 chq.jvb\ndir fmgsql\n$ cd fmgsql\n$ ls\ndir dbnsfp\ndir vvp\n$ cd dbnsfp\n$ ls\n51021 crlq.lrj\n186829 dhcrzvbr.wmn\n16232 fvhn.fqm\n54150 qpbqqj.rpg\n$ cd ..\n$ cd vvp\n$ ls\n179105 rrcsndz.tzp\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd jssnn\n$ ls\ndir bphfqs\ndir dbnsfp\ndir pcccp\ndir snr\ndir zjbvwsnv\n$ cd bphfqs\n$ ls\n110077 dhcrzvbr.wmn\n$ cd ..\n$ cd dbnsfp\n$ ls\ndir hgvh\ndir jtqdcmsz\n154197 rrcsndz.tzp\n$ cd hgvh\n$ ls\ndir qjnbg\n$ cd qjnbg\n$ ls\ndir bqzfpr\n$ cd bqzfpr\n$ ls\n124394 wjsbsp\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd jtqdcmsz\n$ ls\n275597 dbnsfp.fpg\n$ cd ..\n$ cd ..\n$ cd pcccp\n$ ls\ndir cnbd\n85621 cqzvwl\ndir dbnsfp\n114355 hbhp.cfv\ndir mcgq\ndir pcccp\ndir qpbqqj\n224038 rrcsndz.tzp\ndir zcsm\n27570 zjbvwsnv.fjt\n$ cd cnbd\n$ ls\ndir jrbz\ndir pphv\n$ cd jrbz\n$ ls\ndir dwvlwfq\n$ cd dwvlwfq\n$ ls\n32237 fwclr.rnb\n$ cd ..\n$ cd ..\n$ cd pphv\n$ ls\n180370 dhcrzvbr.wmn\n50154 dzvwdwl.gbt\n123965 mlsv.hlw\n163116 wnhtwr.mwl\n$ cd ..\n$ cd ..\n$ cd dbnsfp\n$ ls\n252181 btv.mpv\ndir hwncj\ndir pcccp\n$ cd hwncj\n$ ls\n51410 jbd.fcm\n$ cd ..\n$ cd pcccp\n$ ls\n258123 chq.jvb\n$ cd ..\n$ cd ..\n$ cd mcgq\n$ ls\n206506 qpbqqj.bbb\n$ cd ..\n$ cd pcccp\n$ ls\n193219 ddhtnql.hmb\n134114 hjbpzqzb.rwn\n108927 lznndn.nqd\n73241 ncdrv\n$ cd ..\n$ cd qpbqqj\n$ ls\ndir crdt\ndir tgchdnc\n$ cd crdt\n$ ls\n205710 chq.jvb\n$ cd ..\n$ cd tgchdnc\n$ ls\ndir bdw\ndir dpl\ndir jssnn\ndir pcccp\ndir plpzbm\n$ cd bdw\n$ ls\n211300 dbnsfp.tjm\n$ cd ..\n$ cd dpl\n$ ls\n287744 rsbjqwm\n$ cd ..\n$ cd jssnn\n$ ls\ndir jssnn\n$ cd jssnn\n$ ls\n9644 hmjhshg.vbv\n$ cd ..\n$ cd ..\n$ cd pcccp\n$ ls\ndir jssnn\n85888 pcccp.hdj\ndir qpbqqj\ndir rmscmwtv\n$ cd jssnn\n$ ls\n129698 crlq.lrj\n7499 dhcrzvbr.wmn\n283607 qpbqqj.djr\n234874 wqnrhll\n$ cd ..\n$ cd qpbqqj\n$ ls\n184229 qqpb.ftd\n$ cd ..\n$ cd rmscmwtv\n$ ls\n188048 dhcrzvbr.wmn\ndir jwtpgbnt\n$ cd jwtpgbnt\n$ ls\n209946 hgg\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd plpzbm\n$ ls\n32627 tlb.qmc\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd zcsm\n$ ls\ndir lczflft\ndir zjbvwsnv\ndir zmh\n$ cd lczflft\n$ ls\n40043 dzgnvlw.scr\ndir lrnb\n$ cd lrnb\n$ ls\n109881 bjpbs\ndir jssnn\n46901 npmw\n$ cd jssnn\n$ ls\n9216 sgrp\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd zjbvwsnv\n$ ls\n214676 jssnn.hgn\n$ cd ..\n$ cd zmh\n$ ls\ndir jdt\ndir rggpltr\n$ cd jdt\n$ ls\n147387 jhhsv\n90052 jssnn.wns\n53105 qpbqqj.dzq\n$ cd ..\n$ cd rggpltr\n$ ls\n121454 dbnsfp.dzt\ndir gcc\n$ cd gcc\n$ ls\ndir zbqwl\ndir zjbvwsnv\n$ cd zbqwl\n$ ls\n260297 pcccp.jrw\n$ cd ..\n$ cd zjbvwsnv\n$ ls\n248709 pcccp.tph\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd snr\n$ ls\n152569 chq.jvb\n1437 crlq.lrj\n$ cd ..\n$ cd zjbvwsnv\n$ ls\ndir cqhb\n53235 ghhtl.bhv\n199640 npcfdw\n136346 qpbqqj.lmv\ndir zjbvwsnv\n$ cd cqhb\n$ ls\n24712 sqqf\n$ cd ..\n$ cd zjbvwsnv\n$ ls\ndir gzqg\ndir hfbfvn\ndir srsphr\ndir vgvdcvc\n$ cd gzqg\n$ ls\ndir jjw\n$ cd jjw\n$ ls\n240052 zdcjjz.pmg\n$ cd ..\n$ cd ..\n$ cd hfbfvn\n$ ls\n278190 bfgndw.pvf\n$ cd ..\n$ cd srsphr\n$ ls\n42591 zjbvwsnv.hgh\n$ cd ..\n$ cd vgvdcvc\n$ ls\n120322 rrcsndz.tzp\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd lbrmb\n$ ls\ndir bjhpdj\n42241 crlq.lrj\ndir dbnsfp\n244610 dhcrzvbr.wmn\ndir hppb\ndir mcnzs\ndir npntsr\n13625 tpjpcsgp.dlz\n219424 vvpbt.zvf\ndir zjbvwsnv\n191467 zjbvwsnv.htn\n$ cd bjhpdj\n$ ls\ndir bqjvst\n204722 dbnsfp\ndir dhltrqqq\n226082 dmdqcjp\ndir fcqwgzp\ndir jssnn\n6453 jssnn.ndh\n23799 jssnn.zqn\ndir nwglfhpl\ndir pcccp\ndir pdnj\n269246 shzqns.nws\ndir sjstqlcb\ndir zssln\n$ cd bqjvst\n$ ls\n202793 dbnsfp.pjj\n259783 jssnn\ndir rbvbhnvs\n30683 rvddnjmb.tlz\ndir tzhslnv\n$ cd rbvbhnvs\n$ ls\n86934 vrtrf.htt\n$ cd ..\n$ cd tzhslnv\n$ ls\n76278 mghcwdlr.tsc\n$ cd ..\n$ cd ..\n$ cd dhltrqqq\n$ ls\ndir mfd\ndir pcccp\ndir smmb\n251164 wsdnsgtt.lhr\n191876 zvr.bbz\n$ cd mfd\n$ ls\n51017 crlq.lrj\n99213 rjtbnnnq.hgd\n$ cd ..\n$ cd pcccp\n$ ls\n160487 dhcrzvbr.wmn\ndir nhdrnthj\ndir qpbqqj\n$ cd nhdrnthj\n$ ls\n181291 bbn.wtm\n186551 fnw.tnn\n23622 rrcsndz.tzp\ndir zjbvwsnv\n$ cd zjbvwsnv\n$ ls\n227547 dhcrzvbr.wmn\n$ cd ..\n$ cd ..\n$ cd qpbqqj\n$ ls\n212353 crlq.lrj\n170195 dhcrzvbr.wmn\ndir ttvp\n$ cd ttvp\n$ ls\n185994 tgjcfgjv\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd smmb\n$ ls\ndir dbnsfp\n85354 dbnsfp.zpn\n80665 dfmmjbm.rnr\n135989 dhcrzvbr.wmn\n93718 lrbzr.nfs\ndir mjpfnfns\ndir nsdpfnhb\ndir pmnssvd\n32270 qpbqqj.vtd\n$ cd dbnsfp\n$ ls\n31796 gzs.rgv\n64506 vbjncw.bpz\n181659 vjlfrdp.tqh\n$ cd ..\n$ cd mjpfnfns\n$ ls\n231611 chq.jvb\n17518 cmnlrzq.hvh\n144795 dbnsfp\n162194 jssnn.wjz\n29305 vdqnlw.fzf\n$ cd ..\n$ cd nsdpfnhb\n$ ls\n281844 chq.jvb\n$ cd ..\n$ cd pmnssvd\n$ ls\n165816 dfvl.czb\n144561 gbn\n150785 lnzdwrmb\n111214 rrcsndz.tzp\n164156 tzgdb.hht\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd fcqwgzp\n$ ls\n199161 dhcrzvbr.wmn\n34251 rrcsndz.tzp\n198345 vjlfrdp.tqh\n167001 zjbvwsnv.bsd\n$ cd ..\n$ cd jssnn\n$ ls\ndir ccblfvl\n103180 dhcrzvbr.wmn\ndir prw\ndir tzqfn\ndir zjbvwsnv\n166467 zjbvwsnv.tdt\n$ cd ccblfvl\n$ ls\n159752 crlq.lrj\n20805 jssnn.dvb\n243040 lct.zll\n27492 qbh\n27174 vjlfrdp.tqh\ndir zvfwq\n$ cd zvfwq\n$ ls\n135126 chq.jvb\n41664 gphw.vzd\ndir hmrdghbr\ndir jssnn\ndir qzzb\ndir tmdlcv\n$ cd hmrdghbr\n$ ls\ndir jvgpwrbs\n$ cd jvgpwrbs\n$ ls\ndir wzdv\n$ cd wzdv\n$ ls\n26834 qpbqqj.njf\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd jssnn\n$ ls\n90199 jqqmqddf.qnz\n$ cd ..\n$ cd qzzb\n$ ls\ndir mgpql\ndir src\ndir zvdgc\n$ cd mgpql\n$ ls\n141852 qpbqqj\n$ cd ..\n$ cd src\n$ ls\n204425 lqmcbndm.jrj\n75571 qsbrsv.jcm\n$ cd ..\n$ cd zvdgc\n$ ls\n268742 ffjmrmmz.lhg\n18385 rvmp.hjv\n$ cd ..\n$ cd ..\n$ cd tmdlcv\n$ ls\n182587 sfwvjrj.mzl\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd prw\n$ ls\n207429 dbnsfp.rqf\ndir ptgn\ndir pzgpqp\n252902 rbt\n169694 trg.rsh\n$ cd ptgn\n$ ls\ndir jssnn\ndir qpbqqj\ndir rpd\n$ cd jssnn\n$ ls\n189316 dbnsfp.bqc\n$ cd ..\n$ cd qpbqqj\n$ ls\n167937 zjbvwsnv.bhz\n$ cd ..\n$ cd rpd\n$ ls\n8775 crlq.lrj\n$ cd ..\n$ cd ..\n$ cd pzgpqp\n$ ls\ndir pcccp\n$ cd pcccp\n$ ls\n51496 pcccp\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd tzqfn\n$ ls\ndir cbpfvdp\n285700 crlq.lrj\n7426 dbnsfp.fsd\ndir gdl\n141367 jssnn.hmw\n184482 sczphnp.vnc\n126288 vjlfrdp.tqh\ndir wndpdj\n$ cd cbpfvdp\n$ ls\ndir cvfr\ndir qpbqqj\n$ cd cvfr\n$ ls\ndir jfrnvts\ndir qpbqqj\n$ cd jfrnvts\n$ ls\ndir vwdn\n$ cd vwdn\n$ ls\n236936 vjlfrdp.tqh\n$ cd ..\n$ cd ..\n$ cd qpbqqj\n$ ls\n254275 bqd\n$ cd ..\n$ cd ..\n$ cd qpbqqj\n$ ls\ndir jssnn\n201960 qpbqqj\n$ cd jssnn\n$ ls\n131127 jssnn\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd gdl\n$ ls\n225146 hsgzmtp.wcs\n204436 jssnn.lhh\n64007 mjzjgfg.jsb\n$ cd ..\n$ cd wndpdj\n$ ls\n245412 bvcq\n211386 dbnsfp.tqd\n186962 fql.mww\ndir hlmhtfz\n117446 smvjvcn.lcp\n$ cd hlmhtfz\n$ ls\n150152 lrdhbq.rvm\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd zjbvwsnv\n$ ls\n179703 fvmbz\n87552 qtz.ccw\n129764 rrcsndz.tzp\n$ cd ..\n$ cd ..\n$ cd nwglfhpl\n$ ls\n66039 crlq.lrj\ndir cwq\ndir dlgrsw\n267814 frhlttn.nmd\ndir hmprt\ndir qpbqqj\ndir wnfzznfh\n$ cd cwq\n$ ls\n77655 cpjnwzh\ndir pcccp\ndir zjbvwsnv\ndir zzhjfmnr\n$ cd pcccp\n$ ls\ndir pcccp\n$ cd pcccp\n$ ls\n245309 bggzbrg.flf\n$ cd ..\n$ cd ..\n$ cd zjbvwsnv\n$ ls\n196915 gnmfb.dzq\ndir ngqbdqp\n$ cd ngqbdqp\n$ ls\n355 rrcsndz.tzp\n$ cd ..\n$ cd ..\n$ cd zzhjfmnr\n$ ls\ndir dbnsfp\n$ cd dbnsfp\n$ ls\n223184 chq.jvb\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd dlgrsw\n$ ls\n181906 chq.jvb\n5636 dbnsfp\n219889 jbr.slc\ndir zrntbl\n$ cd zrntbl\n$ ls\n61864 brnpgpwt\n138980 qpbqqj\n$ cd ..\n$ cd ..\n$ cd hmprt\n$ ls\n90249 dbnsfp.mbd\n$ cd ..\n$ cd qpbqqj\n$ ls\n290377 crlq.lrj\n$ cd ..\n$ cd wnfzznfh\n$ ls\n83022 hclmps\n64095 zhm\n$ cd ..\n$ cd ..\n$ cd pcccp\n$ ls\ndir rdzntr\ndir rvccq\n$ cd rdzntr\n$ ls\n239028 rrcsndz.tzp\n$ cd ..\n$ cd rvccq\n$ ls\n22746 chq.jvb\n288752 jjvppq.swt\ndir msgwsnjq\ndir pggz\n228469 vjlfrdp.tqh\n$ cd msgwsnjq\n$ ls\n102522 lwgqc.mhv\n25239 ndm.llf\n$ cd ..\n$ cd pggz\n$ ls\ndir cnjqsqj\n$ cd cnjqsqj\n$ ls\n229407 shpnq\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd pdnj\n$ ls\n193069 rwnhgttz.pvp\n$ cd ..\n$ cd sjstqlcb\n$ ls\n263295 chq.jvb\n224091 jss.wtr\n$ cd ..\n$ cd zssln\n$ ls\n5859 ncdlcr.dll\n$ cd ..\n$ cd ..\n$ cd dbnsfp\n$ ls\n271252 dhcrzvbr.wmn\n$ cd ..\n$ cd hppb\n$ ls\n259968 jssnn\n81292 qpqqb.clj\n$ cd ..\n$ cd mcnzs\n$ ls\n170903 crlq.lrj\n59482 dhcrzvbr.wmn\ndir dqzwzbgm\ndir gnrztn\n286736 jssnn.jcm\n32791 phqsgl\ndir pzjnrwt\n197323 vjlfrdp.tqh\ndir wvnwbpct\n$ cd dqzwzbgm\n$ ls\n78575 qpbqqj\n251546 qpbqqj.slb\n$ cd ..\n$ cd gnrztn\n$ ls\n158603 hdnwmd.rhj\ndir nbfdtwzr\n178239 ptnchzpg\n40517 rrcsndz.tzp\ndir smvb\n198007 vjlfrdp.tqh\n$ cd nbfdtwzr\n$ ls\n200354 crlq.lrj\n$ cd ..\n$ cd smvb\n$ ls\n163921 zjbvwsnv.brz\n$ cd ..\n$ cd ..\n$ cd pzjnrwt\n$ ls\n33891 lwrll\n259646 pcccp.sfn\n106835 pqfzthjq\n189673 rrcsndz.tzp\n$ cd ..\n$ cd wvnwbpct\n$ ls\n234188 dhcrzvbr.wmn\ndir gmtpsgv\n86379 jssnn\n146663 sfpmdbbd.jvt\n25795 vjlfrdp.tqh\n$ cd gmtpsgv\n$ ls\n18642 chq.jvb\n3046 cznlwtw\n26335 ddgpngrc\n116455 vnnls.hsg\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd npntsr\n$ ls\ndir cccjdcvb\n206657 chq.jvb\n280518 crlq.lrj\ndir dbnsfp\ndir jphnn\ndir jssnn\ndir mpl\n195193 rrcsndz.tzp\ndir rztc\ndir znwp\n$ cd cccjdcvb\n$ ls\n192965 mcr.sfq\n$ cd ..\n$ cd dbnsfp\n$ ls\ndir gfns\n173317 jssnn.tjq\ndir mgr\n68817 mvwcwfcr.zmz\ndir pqfht\n108571 swfl.dtj\n10398 tvvvv\ndir vzg\n174361 zjbvwsnv\n$ cd gfns\n$ ls\n203999 zjbvwsnv.hfg\n$ cd ..\n$ cd mgr\n$ ls\ndir zjbvwsnv\n$ cd zjbvwsnv\n$ ls\n26871 tqlgcf.jrn\n$ cd ..\n$ cd ..\n$ cd pqfht\n$ ls\n199590 clpvscl.rlm\ndir dwlhv\ndir vhzfzhrb\n$ cd dwlhv\n$ ls\n130761 qpbqqj\n242752 rrcsndz.tzp\n$ cd ..\n$ cd vhzfzhrb\n$ ls\ndir njdgcbvm\n$ cd njdgcbvm\n$ ls\ndir snjfqg\n$ cd snjfqg\n$ ls\ndir qpwh\n$ cd qpwh\n$ ls\n153353 qsjpj\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd vzg\n$ ls\ndir pcccp\n$ cd pcccp\n$ ls\ndir jfbtl\n$ cd jfbtl\n$ ls\n209199 dbnsfp\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd jphnn\n$ ls\n52305 crlq.lrj\n193480 gmms.whz\n59354 nmq.dww\n64638 qpbqqj\n47072 rrcsndz.tzp\n$ cd ..\n$ cd jssnn\n$ ls\n69168 crlq.lrj\n1549 dhcrzvbr.wmn\n219596 hdmczg.lmm\n108063 jssnn\n24327 vjlfrdp.tqh\ndir zjbvwsnv\n$ cd zjbvwsnv\n$ ls\n189952 chq.jvb\n$ cd ..\n$ cd ..\n$ cd mpl\n$ ls\n144856 bqrrzm\n249487 crlq.lrj\ndir ffqgpgfg\n93632 flqwtn.nsz\ndir mwpcqr\n195910 pdqwn.lcg\n$ cd ffqgpgfg\n$ ls\n66459 dbnsfp\n200500 lcmt.zmz\n207093 qpbqqj\n77042 vjlfrdp.tqh\n57109 wwzv.hbn\n$ cd ..\n$ cd mwpcqr\n$ ls\ndir zjbvwsnv\n$ cd zjbvwsnv\n$ ls\n166393 vjlfrdp.tqh\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd rztc\n$ ls\n57788 chq.jvb\n$ cd ..\n$ cd znwp\n$ ls\n164627 chq.jvb\n$ cd ..\n$ cd ..\n$ cd zjbvwsnv\n$ ls\ndir dgrrl\n71529 jssnn\n198617 pcccp.qqh\ndir phggn\n56842 zjbvwsnv.vqd\n$ cd dgrrl\n$ ls\ndir czm\ndir fhhlbdlz\ndir gstjw\ndir qpbqqj\ndir stgb\n$ cd czm\n$ ls\ndir jssnn\n$ cd jssnn\n$ ls\n162335 chq.jvb\n30099 mfdgdw\n96389 pcdsd.wmw\n251423 tmz.lcb\n205979 vpltdt.gtv\n$ cd ..\n$ cd ..\n$ cd fhhlbdlz\n$ ls\ndir qpbqqj\ndir vdjs\ndir zgz\n$ cd qpbqqj\n$ ls\n285561 chq.jvb\n263924 lbqcfdrs\n138854 pcccp.dtn\n$ cd ..\n$ cd vdjs\n$ ls\n32688 chq.jvb\n223233 tbn.blt\n$ cd ..\n$ cd zgz\n$ ls\n92804 bqltmv.wzb\n$ cd ..\n$ cd ..\n$ cd gstjw\n$ ls\n151784 fvfszzzn.cbh\n$ cd ..\n$ cd qpbqqj\n$ ls\ndir blztqf\ndir plgnh\n$ cd blztqf\n$ ls\n195097 wlvmtz\n$ cd ..\n$ cd plgnh\n$ ls\ndir dbnsfp\n246221 dhcrzvbr.wmn\n271121 jhwmmzls.mhw\n170162 pcccp.dpp\n37872 qpbqqj\n$ cd dbnsfp\n$ ls\ndir dhpnr\n$ cd dhpnr\n$ ls\n152837 pcccp\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd stgb\n$ ls\n248436 vjlfrdp.tqh\n$ cd ..\n$ cd ..\n$ cd phggn\n$ ls\n284602 dhcrzvbr.wmn\ndir lctr\ndir rjmc\n66651 rrcsndz.tzp\n117525 vth.fgw\n156877 wcqnjzbq.dgf\n7803 zpsrzclh.bzw\n$ cd lctr\n$ ls\n212339 jssnn.whp\ndir jzhcqb\n99974 pcccp.zhs\n111354 pmc\n104899 vjlfrdp.tqh\n93496 zhwmbw\n$ cd jzhcqb\n$ ls\ndir zjbvwsnv\n$ cd zjbvwsnv\n$ ls\n146807 rbrg\n$ cd ..\n$ cd ..\n$ cd ..\n$ cd rjmc\n$ ls\ndir fvbmsc\n139747 glwmr.lrg\ndir gvnnz\n102023 tbj.qmz\ndir vsztsjfh\n$ cd fvbmsc\n$ ls\n136838 vpvbz.qtw\n$ cd ..\n$ cd gvnnz\n$ ls\n95498 zjbvwsnv\n$ cd ..\n$ cd vsztsjfh\n$ ls\n215479 ffwlcrwb"

func Part1() int {
	fs := make(map[string][]string, 0)
	currentDir := ""

	for _, line := range strings.Split(input, "\n") {
		isCmd := line[0] == '$'
		var name, arg string
		if isCmd {
			_, _ = fmt.Sscanf(line, "$ %s %s", &name, &arg)
			switch name {
			case "cd":
				{
					if arg == "/" {
						currentDir = arg
					} else if arg == ".." {
						if currentDir == "/" {
							continue
						}
						p := strings.Split(currentDir, "/")
						currentDir = strings.Join(p[:len(p)-2], "/") + "/"
					} else {
						currentDir += fmt.Sprintf("%s/", arg)
					}

					if _, ok := fs[currentDir]; !ok {
						fs[currentDir] = make([]string, 0)
					}
				}
			}
		} else {
			_, _ = fmt.Sscanf(line, "%s %s", &name, &arg)
			if val, ok := fs[currentDir]; ok {
				fs[currentDir] = append(val, line)
			}
		}
	}
	var total int
	for dir, _ := range fs {
		if ds := getSize(dir, fs); ds <= 100_000 {
			total += ds
		}
	}
	return total
}

func getSize(dir string, fs map[string][]string) int {
	contents := fs[dir]
	var size int

	for _, file := range contents {
		var s, name string
		_, _ = fmt.Sscanf(file, "%v %s", &s, &name)
		if s == "dir" {
			size += getSize(dir+name+"/", fs)
		} else {
			ss, _ := strconv.Atoi(s)
			size += ss
		}
	}
	return size
}

func Part2() int {
	fs := make(map[string][]string, 0)
	currentDir := ""

	for _, line := range strings.Split(input, "\n") {
		isCmd := line[0] == '$'
		var name, arg string
		if isCmd {
			_, _ = fmt.Sscanf(line, "$ %s %s", &name, &arg)
			switch name {
			case "cd":
				{
					if arg == "/" {
						currentDir = arg
					} else if arg == ".." {
						if currentDir == "/" {
							continue
						}
						p := strings.Split(currentDir, "/")
						currentDir = strings.Join(p[:len(p)-2], "/") + "/"
					} else {
						currentDir += fmt.Sprintf("%s/", arg)
					}

					if _, ok := fs[currentDir]; !ok {
						fs[currentDir] = make([]string, 0)
					}
				}
			}
		} else {
			_, _ = fmt.Sscanf(line, "%s %s", &name, &arg)
			if val, ok := fs[currentDir]; ok {
				fs[currentDir] = append(val, line)
			}
		}
	}

	// Part 1 changes start here
	used := getSize("/", fs)
	available := 70_000_000 - used
	needed := 30_000_000 - available
	var best = used
	for dir, _ := range fs {
		size := getSize(dir, fs)
		if size >= needed {
			if size < best {
				best = size
			}

		}
	}
	return best
}
