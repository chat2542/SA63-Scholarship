/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Playlist Vidoe
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntScholarshipinformation,
    EntScholarshipinformationFromJSON,
    EntScholarshipinformationFromJSONTyped,
    EntScholarshipinformationToJSON,
    EntScholarshiptype,
    EntScholarshiptypeFromJSON,
    EntScholarshiptypeFromJSONTyped,
    EntScholarshiptypeToJSON,
    EntSemester,
    EntSemesterFromJSON,
    EntSemesterFromJSONTyped,
    EntSemesterToJSON,
    EntUser,
    EntUserFromJSON,
    EntUserFromJSONTyped,
    EntUserToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntScholarshipRequestEdges
 */
export interface EntScholarshipRequestEdges {
    /**
     * 
     * @type {EntScholarshipinformation}
     * @memberof EntScholarshipRequestEdges
     */
    scholarshipinformation?: EntScholarshipinformation;
    /**
     * 
     * @type {EntScholarshiptype}
     * @memberof EntScholarshipRequestEdges
     */
    scholarshiptype?: EntScholarshiptype;
    /**
     * 
     * @type {EntSemester}
     * @memberof EntScholarshipRequestEdges
     */
    semester?: EntSemester;
    /**
     * 
     * @type {EntUser}
     * @memberof EntScholarshipRequestEdges
     */
    user?: EntUser;
}

export function EntScholarshipRequestEdgesFromJSON(json: any): EntScholarshipRequestEdges {
    return EntScholarshipRequestEdgesFromJSONTyped(json, false);
}

export function EntScholarshipRequestEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntScholarshipRequestEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'scholarshipinformation': !exists(json, 'scholarshipinformation') ? undefined : EntScholarshipinformationFromJSON(json['scholarshipinformation']),
        'scholarshiptype': !exists(json, 'scholarshiptype') ? undefined : EntScholarshiptypeFromJSON(json['scholarshiptype']),
        'semester': !exists(json, 'semester') ? undefined : EntSemesterFromJSON(json['semester']),
        'user': !exists(json, 'user') ? undefined : EntUserFromJSON(json['user']),
    };
}

export function EntScholarshipRequestEdgesToJSON(value?: EntScholarshipRequestEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'scholarshipinformation': EntScholarshipinformationToJSON(value.scholarshipinformation),
        'scholarshiptype': EntScholarshiptypeToJSON(value.scholarshiptype),
        'semester': EntSemesterToJSON(value.semester),
        'user': EntUserToJSON(value.user),
    };
}


