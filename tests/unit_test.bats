#!/usr/bin/env bats

@test "test error no argument" {
    run bash -c "echo $? | ./207demography"
    [ "$status" -eq 84 ]
}

@test "test error invalid country" {
    run bash -c "echo $? | ./207demography MANOU"
    [ "$status" -eq 84 ]
}

@test "test NaN" {
    run bash -c "echo NaN | ./207demography 1"
    [ "$status" -eq 84 ]
}

@test "test error invalid countries" {
    run bash -c "echo $? | ./207demography MANOU"
    [ "$status" -eq 84 ]
}

@test "test EUU country" {
    run bash -c "./207demography EUU > EUU.txt"
    run bash -c "diff filetest/pdfEUU.txt EUU.txt"
    [ "$status" -eq 0 ]
    run bash -c "rm EUU.txt"
}

@test "test BRA BOL PER country" {
    run bash -c "./207demography BRA BOL PER > BRA_BRO_PER.txt"
    run bash -c "diff filetest/pdfBRA_BRO_PER.txt BRA_BRO_PER.txt"
    [ "$status" -eq 0 ]
    run bash -c "rm BRA_BRO_PER.txt"
}

@test "test OED members" {
    run bash -c "./207demography OED > OED.txt"
    run bash -c "diff filetest/pdfOED.txt OED.txt"
    [ "$status" -eq 0 ]
    run bash -c "rm OED.txt"
}

@test "test WLD" {
    run bash -c "./207demography WLD > WLD.txt"
    run bash -c "diff filetest/pdfWLD.txt WLD.txt"
    [ "$status" -eq 0 ]
    run bash -c "rm WLD.txt"
}

@test "test Not classified" {
    run bash -c "./207demography INX > INX.txt"
    run bash -c "diff filetest/pdfINX.txt INX.txt"
    [ "$status" -eq 0 ]
    run bash -c "rm INX.txt"
}

@test "test african country" {
    run bash -c "./207demography MDG TUN MAR EGY > fav.txt"
    run bash -c "diff filetest/pdfFAV.txt fav.txt"
    [ "$status" -eq 0 ]
    run bash -c "rm fav.txt"
}