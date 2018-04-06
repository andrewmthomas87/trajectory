package main

import (
	"./trajectory"
	"fmt"
	"math"
	"io/ioutil"
)

func main() {
	robot := &trajectory.SkidSteerRobot{
		Acceleration:        6,
		Deceleration:        6,
		MaxVelocity:         6,
		AngularAcceleration: math.Pi,
	}

	writeProfileFileFor("SwitchLeftFromCenter", robot, 500, 90, 90, trajectory.Point{0, 0}, trajectory.Point{3, 4})
	writeProfileFileFor("SwitchLeftFromCenter", robot, 500, 90, 90, trajectory.Point{0, 0}, trajectory.Point{3, 4})
	writeProfileFileFor("SwitchRightFromCenter", robot, 500, 90, 90, trajectory.Point{0, 0}, trajectory.Point{-4, 4})
}

const formatString = `package org.usfirst.frc.team1619.robot.trajectories;

import org.usfirst.frc.team1619.robot.framework.trajectory.TrajectoryData;

public class %v extends TrajectoryData {
	public %v() {
		super(
			new double[]{%v},
			new double[]{%v},
			new double[]{%v},
			new double[]{%v}
		)
	}
}
`

func writeProfileFileFor(className string, robot *trajectory.SkidSteerRobot, resolution int, startHeadingDegrees, endHeadingDegrees float64, points ...trajectory.Point) {
	spline := trajectory.SplineFor(startHeadingDegrees, endHeadingDegrees, points...)
	profile := trajectory.NewVelocityProfile(robot, spline)

	profile.Calculate(resolution)

	timeStr := ""
	distanceStr := ""
	velocityStr := ""
	headingStr := ""
	for i := 0; i <= resolution; i++ {
		timeStr += fmt.Sprintf("%G,", profile.TimeValues[i])
		distanceStr += fmt.Sprintf("%G,", profile.DistanceValues[i])
		velocityStr += fmt.Sprintf("%G,", profile.VelocityValues[i])
		headingStr += fmt.Sprintf("%G,", profile.HeadingValues[i])
	}

	classStr := fmt.Sprintf(formatString, className, className, timeStr, distanceStr, velocityStr, headingStr)
	ioutil.WriteFile(fmt.Sprintf("./trajectories/%v.java", className), []byte(classStr), 777)
}
